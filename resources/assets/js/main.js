const toggleButton = document.getElementById("toggle-dark-mode");
const rootElement = document.documentElement;

// Function to apply the dark mode based on saved preference
function applySavedTheme() {
  const savedTheme = localStorage.getItem("theme");
  if (savedTheme === "dark") {
    rootElement.setAttribute("data-theme", "dark");
  }
}

// Initialize the theme on page load
applySavedTheme();

toggleButton.addEventListener("click", () => {
  if (rootElement.hasAttribute("data-theme")) {
    rootElement.removeAttribute("data-theme");
    localStorage.removeItem("theme");
  } else {
    rootElement.setAttribute("data-theme", "dark");
    localStorage.setItem("theme", "dark");
  }
});
