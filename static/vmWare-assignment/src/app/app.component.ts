import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { FileUploadService } from './file-upload.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = 'vmWare-assignment';
  allAlbumNames: String[] = []
  allImageNames: String[] = []
  isAlbumChosen: Boolean = false
  chosenAlbum:String = ""
  newAlbumForm: FormGroup
  file: File = null;
  baseUrl = "http://ip172-18-0-28-ccjmdqld48eg00a85oq0-80.direct.labs.play-with-docker.com"

  constructor(private route: ActivatedRoute, private router: Router,private fileUploadService: FileUploadService) {}

  ngOnInit() {
    this.newAlbumForm = new FormGroup({
      newAlbumName: new FormControl("")
    })

    let url = `${this.baseUrl}/getAllAlbums`;
      fetch(url)
        .then(res => res.json())
        .then(response => {
              this.allAlbumNames = response
              console.log(this.allAlbumNames)
          });
  }


  submitAlbumName(){
    let name = this.newAlbumForm.value.newAlbumName;
    let url = `${this.baseUrl}/createAlbum/${name}`;
      fetch(url)
        .then(res => res.json())
        .then(response => {
              this.allAlbumNames = response
              console.log(this.allAlbumNames)
          });
          this.isAlbumChosen =false
          this.chosenAlbum = ""
  }

  onChange(event) {
    this.file = event.target.files[0];
  }

  onUpload(){
    console.log(this.file);
    this.fileUploadService.upload(this.file,this.chosenAlbum,this.baseUrl).subscribe(
        (event: any) => {
          console.log(event)
        }
    );
  }

  deleteAlbum(albumName){
    let url = `${this.baseUrl}/deleteAlbum/${albumName}`;
    fetch(url)
      .then(res => res.json())
      .then(response => {
            this.allAlbumNames = response
            console.log(this.allAlbumNames)
        });
        this.isAlbumChosen =false
        this.chosenAlbum = ""
  }

  getAllImages(albumName){
    let url = `${this.baseUrl}/getAllImages/${albumName}`;
    this.isAlbumChosen = true
    this.chosenAlbum = albumName
    fetch(url)
      .then(res => res.json())
      .then(response => {
            this.allImageNames = response
            console.log(this.allImageNames)
        });
  }

  deleteImage(img,chosenAlbum){
    let url = `${this.baseUrl}/deleteImageFromAlbum/${chosenAlbum}/${img}`;
    this.isAlbumChosen = true
    this.chosenAlbum = chosenAlbum
    fetch(url)
      .then(res => res.json())
      .then(response => {
            this.allImageNames = response
            console.log(this.allImageNames)
        });
        this.isAlbumChosen = true
        this.chosenAlbum = chosenAlbum
  }




}
