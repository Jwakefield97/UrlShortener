
function getSearchedLinks(searchTerm, linkMap) {
	let resultArr = [];
	for(let prop in linkMap) {
		if(prop === searchTerm || prop.includes(searchTerm) || 
					linkMap[prop] === searchTerm  || linkMap[prop].includes(searchTerm)) {
			resultArr.push([linkMap[prop],prop])
		}
	}
	return resultArr;
}

//just a cheap way of making sure there isn't any js in the urls
//TODO: fix this 
function getEscapedString(str) {
	str = str.replace(/</g, "");
	str = str.replace(/>/g, "");
	return str;
}

function populateSearchResults(searchResults) {
	let resultsUl = document.getElementById("searchResults");
	let outputStr = searchResults.map(res => {
		let link1 = getEscapedString(res[0]),
			link2 = getEscapedString(res[1]);

		return `<li class="collection-item"> 
					<b>Original:</b> <a href="${link1}">${link1}</a> 
					<b>Shortened:</b> <a href="${SHORTENED_LINK_BASE}${link2}">${SHORTENED_LINK_BASE}${link2}</a> 
				</li>`
	}).join("");
	resultsUl.innerHTML = outputStr;
}

window.addEventListener('DOMContentLoaded', (event) => {
	let allUrls = {},
		searchInput = document.getElementById("searchInput");

	fetch("/urls")
	.then(result => result.json())
	.then(json => {
		allUrls = json;
		populateSearchResults(getSearchedLinks("", allUrls));
	});

	searchInput.addEventListener("change", (e) => {
		populateSearchResults(getSearchedLinks(searchInput.value, allUrls));
	});
});