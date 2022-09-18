import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class FileUploadService {
  // API url
  baseApiUrl = ""
    
  constructor(private http:HttpClient) { }

  upload(file,albumName,baseUrl):Observable<any> {
      this.baseApiUrl = `${baseUrl}/uploadImage/${albumName}`
      const formData = new FormData(); 
        
      // Store form name as "file" with file data
      formData.append("file", file, file.name);
  
      return this.http.post(this.baseApiUrl, formData)
  }
}
