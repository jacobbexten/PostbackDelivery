# Postback Delivery Mini Project

## **Description**  
A PHP application that ingests http requests (RAW POST DATA) along with a Go application that delivers http responses stored in a log file. Redis is used as a queue between them.  
  
<br />

Requests were tested using Postman. https://www.postman.com/  

<br />

### **Sample Request:**  
(POST) http://{server_ip}/ingest.php  
(RAW POST DATA)  
{  
&nbsp;"endpoint":{  
&nbsp;&nbsp;"method":"GET",  
&nbsp;&nbsp;"url":"http://sample_domain_endpoint.com/data?title={mascot}&image={location}&foo={bar}"  
&nbsp;},  
&nbsp;"data":[  
&nbsp;&nbsp;{  
&nbsp;&nbsp;&nbsp;"mascot":"Gopher",  
&nbsp;&nbsp;&nbsp;"location":"https://blog.golang.org/gopher/gopher.png"  
&nbsp;&nbsp;}  
&nbsp;]  
}  

<br />

### **Sample Response (Postback):**
GET
http://sample_domain_endpoint.com/data?title=Gopher&image=https%3A%2F%2Fblog.golang.org%2Fg
opher%2Fgopher.png&foo=
