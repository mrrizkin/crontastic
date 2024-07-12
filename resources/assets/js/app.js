import "../css/index.css";
import "preline";
import "./icons";

document.addEventListener("DOMContentLoaded", function () {
  const navLinks = document.querySelectorAll(".nav-link");
  for (const navLink of navLinks) {
    if (navLink.href === window.location.href) {
      navLink.classList.add("active");
    }
  }
});
