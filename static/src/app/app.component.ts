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
  apiUrl: String = ""
  baseApiUrl = ""

  constructor(private route: ActivatedRoute, private router: Router,private fileUploadService: FileUploadService) {
    this.baseApiUrl = fileUploadService.baseApiUrl
  }

  ngOnInit() {
    this.newAlbumForm = new FormGroup({
      newAlbumName: new FormControl("")
    })

    this.apiUrl = `/getAllAlbums`;

    this.fileUploadService.fetchData(this.apiUrl).subscribe(
      (response: any) => {
        this.allAlbumNames = response
        console.log(this.allAlbumNames)
      });
  }


  submitAlbumName(){
    let name = this.newAlbumForm.value.newAlbumName;
    this.apiUrl = `/createAlbum/${name}`;
    this.fileUploadService.fetchData(this.apiUrl).subscribe(
      (response: any) => {
        this.allAlbumNames = response
        console.log(this.allAlbumNames)
        this.isAlbumChosen =false
        this.chosenAlbum = ""
      });
  }

  onChange(event) {
    this.file = event.target.files[0];
  }

  onUpload(){
    console.log(this.file);
    this.apiUrl = `/uploadImage/${this.chosenAlbum}`
    this.fileUploadService.upload(this.file,this.apiUrl).subscribe(
        (response: any) => {
          this.allImageNames = response
          console.log(this.allImageNames)
          this.isAlbumChosen = true
          this.chosenAlbum = this.chosenAlbum
        }
    );
  }

  deleteAlbum(albumName){
    this.apiUrl = `/deleteAlbum/${albumName}`;
    this.fileUploadService.fetchData(this.apiUrl).subscribe(
      (response: any) => {
        this.allAlbumNames = response
        console.log(this.allAlbumNames)
        this.isAlbumChosen =false
        this.chosenAlbum = ""
      });
  }

  getAllImages(albumName){
    this.apiUrl = `/getAllImages/${albumName}`;
    this.isAlbumChosen = true
    this.chosenAlbum = albumName

    this.fileUploadService.fetchData(this.apiUrl).subscribe(
      (response: any) => {
        this.allImageNames = response
        console.log(this.allImageNames)
      });
  }

  deleteImage(img,chosenAlbum){
    this.apiUrl = `/deleteImageFromAlbum/${chosenAlbum}/${img}`;
    this.isAlbumChosen = true
    this.chosenAlbum = chosenAlbum

    this.fileUploadService.fetchData(this.apiUrl).subscribe(
      (response: any) => {
        this.allImageNames = response
        console.log(this.allImageNames)
        this.isAlbumChosen = true
        this.chosenAlbum = chosenAlbum
      });
  }




}
