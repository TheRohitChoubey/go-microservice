let allAlbumNames = []
let chosenAlbum
let baseUrl = "http://localhost:80"
allImageNames = []

window.onload = () => {
    loadAlbums()
}
const createNewAlbum = () => {
    let albumName = document.getElementById("newAlbumName").value;
    let url = `${baseUrl}/createAlbum/${albumName}`;
    fetch(url)
        .then(res => res.json())
        .then(response => {
            this.allAlbumNames = response
            console.log(this.allAlbumNames)
            loadAlbums()
        })

}

const loadAlbums = () => {
    let url = `${baseUrl}/getAllAlbums`;
    fetch(url)
        .then(res => res.json())
        .then(response => {
            this.allAlbumNames = response
            console.log(this.allAlbumNames)
            this.chosenAlbum = ""
            let allAlbumHtml = ""
            for (let i = 0; i < this.allAlbumNames.length; i++) {
                allAlbumHtml += `<button class="nav-link active" onclick="getAllImages('${this.allAlbumNames[i]}')"> ${this.allAlbumNames[i]}
                <button class="btn" onclick="deleteAlbum('${this.allAlbumNames[i]}')"><i class="fa fa-trash"></i></button>
            </button>`

            }

            document.getElementById('allAlbums').innerHTML = allAlbumHtml;

        })
}

const deleteAlbum = (albumName) => {
    let url = `${baseUrl}/deleteAlbum/${albumName}`;
    fetch(url)
        .then(res => res.json())
        .then(response => {
            loadAlbums()
        })
}

const getAllImages = (albumName) => {
    let url = `${baseUrl}/getAllImages/${albumName}`;

    fetch(url)
        .then(res => res.json())
        .then(response => {
            this.allImageNames = response
            console.log(this.allImageNames)

            let albumDivHtml = `<div class="col-6">
            <div class="card border-primary">
                <h6 class="card-header h2">Upload a new image to ${albumName} Album</h6>
                <div class="card-body text-success">
                    <div class="form-group">
                        <div class="row">
                            <div class="form-group">
                                <label for="inputEmail4" class="text-white">Choose the image to be uploaded ... </label>
                                <input class="form-control" type="file" id="imageFile" name="imageFile"><br><br>
                                <input type="button" onclick="submitImage('${albumName}')" value="Upload">
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>`

            if (this.allImageNames != null) {
                for (let i = 0; i < this.allImageNames.length; i++) {
                    albumDivHtml += `
            <div class="row">
                <div class="col-md-4">
                    <a class="thumbnail" target="_blank" href="../images/${albumName}/${this.allImageNames[i]}">
                            <img class="img-responsive" src="../images/${albumName}/${this.allImageNames[i]}" alt="${this.allImageNames[i]}"> Click here to open
                        </a> 
                    <button class="btn" onclick="deleteImage('${this.allImageNames[i]}','${albumName}')"><i class="fa fa-trash"></i></button>
                </div>
            </div>`
                }
            }

            document.getElementById('albumDiv').innerHTML = albumDivHtml;


        })
}

const deleteImage = (imageName, albumName) => {
    let url = `${baseUrl}/deleteImageFromAlbum/${albumName}/${imageName}`;
    fetch(url)
        .then(res => res.json())
        .then(response => {
            getAllImages(albumName)
        })
}

const submitImage = (albumName) => {
    let input = document.getElementById("imageFile")
    let data = new FormData()
    data.append('file', input.files[0])

    let url = `${baseUrl}/uploadImage/${albumName}`;

    fetch(url, {
            method: "POST",
            body: data
        })
        .then(response => {
            getAllImages(albumName)
        })

}