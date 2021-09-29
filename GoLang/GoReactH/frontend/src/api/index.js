var connection = "TODO"

function getAlbums() {
    fetch("/api", {method: 'GET'})
    .then(response => { if (response.ok) {console.log("good"); return response.json();} else {console.log("failed here");throw response}})
    .then(data => console.log(data))
    .catch(error => console.log(error))
}

export {getAlbums}