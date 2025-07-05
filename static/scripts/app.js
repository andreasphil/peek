import { useThemeColor } from "@external/andreasphil/design-system@v0.49.0/scripts/utils.js";
import { highlightElement } from "@external/@speed-highlight/core@1.2.7/dist/index.js";

useThemeColor();

// Syntax highlighting ------------------------------------

const codeBlockEls = document.querySelectorAll("pre");
codeBlockEls.forEach((el) => {
  const codeEl = el.querySelector("code[class|=language]");
  const lang = codeEl.className.match(/language-([\w-]+)/)?.[1];

  if (lang) {
    highlightElement(el, lang, "multiline", { hideLineNumbers: true });
  }
});

// Table of Contents --------------------------------------

const tocSelect = document.getElementById("toc-select");
const headingEls = document.querySelectorAll("h1,h2,h3");

const visibleHeadingObserver = new IntersectionObserver((e) => {
  const current = e.find((i) => i.isIntersecting)?.target?.id;
  if (current) tocSelect.value = current;
});

headingEls.forEach((i) => {
  // IDs should be generated my the markdown renderer but sometimes they're not,
  // adding a fallback in case that happens
  if (!i.id) i.id = crypto.randomUUID();

  // Populate TOC select with headings
  const optionEl = document.createElement("option");
  optionEl.value = i.id;
  optionEl.textContent = i.textContent;
  tocSelect.appendChild(optionEl);

  // Observe scroll state
  visibleHeadingObserver.observe(i);
});

tocSelect.addEventListener("change", ({ target }) => {
  const targetHeading = document.getElementById(target.value);
  targetHeading.scrollIntoView({ behavior: "smooth" });
  history.pushState(null, null, `#${target.value}`);
});

document.addEventListener("keyup", (e) => {
  if (e.key === "/") tocSelect.focus();
});

// Sharing ------------------------------------------------

function getContent() {
  const content = document.getElementById("content");
  return { plain: content.innerText, html: content.innerHTML };
}

document.getElementById("share-button").addEventListener("click", () => {
  const { html } = getContent();
  const blob = new Blob([html], { type: "text/html" });
  const file = new File([blob], "shared-content.html", { type: "text/html" });
  navigator.share({ files: [file] });
});

document.getElementById("copy-button").addEventListener("click", () => {
  const { plain, html } = getContent();
  navigator.clipboard.write([
    new ClipboardItem({
      "text/html": new Blob([html], { type: "text/html" }),
      "text/plain": new Blob([plain], { type: "text/plain" }),
    }),
  ]);
});
