var navLinks = document.querySelectorAll("nav a");
for (var i = 0; i < navLinks.length; i++) {
	var link = navLinks[i]
	if (link.getAttribute('href') == window.location.pathname) {
		link.classList.add("live");
		break;
	}
}

// Toggle password visibility for inputs with a matching data-target
var toggles = document.querySelectorAll(".toggle-password");
for (var j = 0; j < toggles.length; j++) {
	var btn = toggles[j];
	var targetId = btn.getAttribute("data-target");
	var input = document.getElementById(targetId);
	if (!input) {
		continue;
	}
	btn.addEventListener("click", (function(button, field) {
		return function () {
			var showing = field.getAttribute("type") === "text";
			field.setAttribute("type", showing ? "password" : "text");
			button.classList.toggle("active", !showing);
			button.setAttribute("aria-label", showing ? "Show password" : "Hide password");
		}
	})(btn, input));
}