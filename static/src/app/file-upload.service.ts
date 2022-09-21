import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class FileUploadService {
  // API url
  baseApiUrl = "http://localhost:80"
    
  constructor(private http:HttpClient) { }

  upload(file,api):Observable<any> {
      let apiUrl = this.baseApiUrl+ api
      const formData = new FormData(); 
        
      // Store form name as "file" with file data
      formData.append("file", file, file.name);
  
      return this.http.post(apiUrl, formData)
  }

  fetchData(api):Observable<any> {
    let apiUrl = this.baseApiUrl+ api
    return this.http.get(apiUrl)
  }


}
