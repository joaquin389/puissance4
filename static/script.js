document.addEventListener("DOMContentLoaded", () => {
  const body = document.body;
  const toggle = document.querySelector(".switch input");

  // Appliquer le thème sauvegardé
  const savedTheme = localStorage.getItem("theme");
  if (savedTheme === "light") {
    body.classList.add("light-mode");
    body.classList.remove("dark-mode");
    toggle.checked = false;
  } else {
    body.classList.add("dark-mode");
    body.classList.remove("light-mode");
    toggle.checked = true;
  }

  // Écoute du changement de thème
  toggle.addEventListener("change", () => {
    if (toggle.checked) {
      body.classList.add("dark-mode");
      body.classList.remove("light-mode");
      localStorage.setItem("theme", "dark");
    } else {
      body.classList.add("light-mode");
      body.classList.remove("dark-mode");
      localStorage.setItem("theme", "light");
    }
  });

  // Restaurer la position de scroll sans animation
  const savedScrollY = localStorage.getItem("scrollY");
  if (savedScrollY !== null) {
    // Attendre que le layout soit prêt
    requestAnimationFrame(() => {
      window.scrollTo(0, parseInt(savedScrollY));
    });
  }
});

// Sauvegarder la position de scroll à chaque scroll
window.addEventListener("scroll", () => {
  localStorage.setItem("scrollY", window.scrollY);
});
