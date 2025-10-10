document.addEventListener("DOMContentLoaded", () => {
  const saved = localStorage.getItem("darkMode");
  const isDark = saved === "true";

  // ✅ Crée une checkbox invisible si elle n'existe pas
  let checkbox = document.getElementById("checkbox");
  if (!checkbox) {
    checkbox = document.createElement("input");
    checkbox.type = "checkbox";
    checkbox.id = "checkbox";
    checkbox.style.display = "none";
    document.body.appendChild(checkbox);
  }

  // ✅ Applique l'état sauvegardé
  checkbox.checked = isDark;

  // ✅ Synchronise avec le switch visible si présent
  const visibleSwitch = document.querySelector('.switch input');
  if (visibleSwitch) {
    visibleSwitch.checked = isDark;

    visibleSwitch.addEventListener("change", () => {
      checkbox.checked = visibleSwitch.checked;
      localStorage.setItem("darkMode", checkbox.checked);
    });
  }

  // ✅ Sauvegarde si la checkbox change (même invisible)
  checkbox.addEventListener("change", () => {
    localStorage.setItem("darkMode", checkbox.checked);
  });
});

// avec cookie
// variable de page en page
