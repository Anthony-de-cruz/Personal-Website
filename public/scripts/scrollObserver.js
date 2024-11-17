const observer = new IntersectionObserver((entries) => {
  entries.forEach((entry) => {
    console.log(entry);
    if (entry.isIntersecting) {
      entry.target.classList.add("scroll-show");
    } else {
      entry.target.classList.remove("scroll-show");
    }
  });
});
function observeHiddenElements() {
  const hiddenElements = document.querySelectorAll(".scroll-hidden");
  hiddenElements.forEach((el) => observer.observe(el));
}

// Reobserve new elements after HTMX swaps
document.body.addEventListener("htmx:afterSwap", () => {
  observeHiddenElements();
});

observeHiddenElements();
