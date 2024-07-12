import "../css/index.css";
import "preline";

import {
  createIcons,
  ArrowRight,
  ArrowRightLeft,
  BellDot,
  BookOpenText,
  BookText,
  CalendarCheck,
  ClipboardList,
  DatabaseBackup,
  Download,
  House,
  ListRestart,
  RefreshCcw,
  RotateCcw,
  ScrollText,
  Settings,
  ShieldCheck,
  ShieldEllipsis,
  Trash2,
  UsersRound,
} from "lucide";

createIcons({
  icons: {
    ArrowRight,
    ArrowRightLeft,
    BellDot,
    BookOpenText,
    BookText,
    CalendarCheck,
    ClipboardList,
    DatabaseBackup,
    Download,
    House,
    ListRestart,
    RefreshCcw,
    RotateCcw,
    ScrollText,
    Settings,
    ShieldCheck,
    ShieldEllipsis,
    Trash2,
    UsersRound,
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
