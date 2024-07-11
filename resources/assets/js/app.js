import "../css/index.css";
import "preline";

import { createIcons, ClipboardList, BookOpen, House, CalendarCheck, ArrowRight, UsersRound, Settings } from "lucide";

createIcons({
  icons: {
    House,
    BookOpen,
    CalendarCheck,
    ClipboardList,
    UsersRound,
    ArrowRight,
    Settings,
  },
});

document.addEventListener("DOMContentLoaded", function () {
  const navLinks = document.querySelectorAll(".nav-link");
  for (const navLink of navLinks) {
    if (navLink.href === window.location.href) {
      navLink.classList.add("active");
    }
  }
});
