import { HttpErrorResponse } from '@angular/common/http';
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
  uploadImageForm: FormGroup
  file: File = null;
  apiUrl: String = ""
  baseApiUrl = ""
  extension = ["jpg","png","jpeg","gif"]

  constructor(private route: ActivatedRoute, private router: Router,private fileUploadService: FileUploadService) {
    this.baseApiUrl = fileUploadService.baseApiUrl
  }

  ngOnInit() {
    this.newAlbumForm = new FormGroup({
      newAlbumName: new FormControl("")
    })

    this.uploadImageForm = new FormGroup({
      fileUploadControl : new FormControl()
    })
    
    this.apiUrl = `/getAllAlbums`;

    this.fileUploadService.fetchData(this.apiUrl).subscribe(
      (response: any) => {
        this.allAlbumNames = response
      });
  }


  submitAlbumName(){
    let name = this.newAlbumForm.value.newAlbumName;
    this.apiUrl = `/createAlbum/${name}`;
    this.fileUploadService.fetchData(this.apiUrl).subscribe(
      (response: any) => {
        this.allAlbumNames = response
        this.isAlbumChosen =false
        this.chosenAlbum = ""
        this.newAlbumForm.reset()
        alert("New Album Created Successfully")
        
      });
  }

  onChange(event) {
    this.file = event.target.files[0];
    let ext = this.file.name.split('.').pop(); 
    let allowed = this.extension.find( p => {
      p == ext
    })
    if(allowed){
      this.uploadImageForm.reset()
      alert("Please select a image file to upload, only jpg, png, jpeg, gif are allowed")
    }
  }

  onUpload(){
    console.log(this.file);
    this.apiUrl = `/uploadImage/${this.chosenAlbum}`
    if(this.file != null){
      this.fileUploadService.upload(this.file,this.apiUrl).subscribe(
          (response: any) => {
            this.isAlbumChosen = true
            this.chosenAlbum = this.chosenAlbum
            alert("File uploaded Successfully")
          }, (err : any) => {
            this.isAlbumChosen = true
            this.chosenAlbum = this.chosenAlbum
            this.getAllImages(this.chosenAlbum)
            alert("File uploaded Successfully")
          }
          
      );
      this.uploadImageForm.reset()
      alert("File upload in progress")
    } else {
      alert("Please select a file to upload")
    }
  }

  deleteAlbum(albumName){
    this.apiUrl = `/deleteAlbum/${albumName}`;
    this.fileUploadService.fetchData(this.apiUrl).subscribe(
      (response: any) => {
        this.allAlbumNames = response
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
      });
  }

  deleteImage(img,chosenAlbum){
    this.apiUrl = `/deleteImageFromAlbum/${chosenAlbum}/${img}`;
    this.isAlbumChosen = true
    this.chosenAlbum = chosenAlbum

    this.fileUploadService.fetchData(this.apiUrl).subscribe(
      (response: any) => {
        this.allImageNames = response
        this.isAlbumChosen = true
        this.chosenAlbum = chosenAlbum
      });
  }




}
