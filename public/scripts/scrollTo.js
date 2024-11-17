document.addEventListener("htmx:afterSwap", (event) => {
  // Scroll to target on swap
  if (event.target.id === "content") {
    document.querySelector("#content").scrollIntoView({
      behavior: "smooth",
      block: "start",
    });
  }
});
