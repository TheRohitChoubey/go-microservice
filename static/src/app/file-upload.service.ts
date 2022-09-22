import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class FileUploadService {
  // API url
  // localhost:80
  // ip172-18-0-52-ccm04koja8q000cnvcd0-80.direct.labs.play-with-docker.com
  baseApiUrl = "http://ip172-18-0-52-ccm04koja8q000cnvcd0-80.direct.labs.play-with-docker.com"
    
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
