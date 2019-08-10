function startLoader(submitBtn, formLoader) {
	submitBtn.classList.add("disabled");
	formLoader.classList.add("show");
	formLoader.classList.remove("hide");
}

function stopLoader(submitBtn, formLoader) {
	submitBtn.classList.remove("disabled");
	formLoader.classList.add("hide");
	formLoader.classList.remove("show");
}

function showNewUrlNode(newUrlLinkWrapper, newUrlLink, newUrl) {
	newUrlLinkWrapper.classList.remove("hide");
	newUrlLink.href = newUrl;
	newUrlLink.innerHTML = ` ${newUrl}`;
}

window.addEventListener('DOMContentLoaded', (event) => {
	let submitButton = document.getElementById("submitNewUrl"),
		urlInput = document.getElementById("newURL"),
		formLoader = document.getElementById("formLoader"),
		newUrlLink = document.getElementById("newUrlLink"),
		newUrlLinkWrapper = document.getElementById("newUrlLinkWrapper");
	
	submitButton.addEventListener("click", (e) => {
		if(urlInput.value.trim() !== "") {
			startLoader(submitButton, formLoader);
			let formData = new FormData();
			formData.append("url", urlInput.value);
			fetch("/new", {
				method: "POST",
				body: formData
			})
			.then(res => res.json())
			.then(response => {
				stopLoader(submitButton, formLoader);
				showNewUrlNode(newUrlLinkWrapper, newUrlLink, response.url);
			})
			.catch(error => {
				stopLoader(submitButton, formLoader);
				console.error(error);
			});
		}
	});
});




