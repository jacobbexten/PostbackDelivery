# Postback Delivery Mini Project

## **Description**  
A [PHP server](ingestion_agent/ingest.php) that ingests http requests (RAW POST DATA) along with a [Go server](delivery_agent/main.go) that delivers http responses stored in a log file. Redis is used as a queue between them.  
  
<br />

Requests were tested using Postman https://www.postman.com/  and a local Go server through [test_endpoint.go](delivery_agent/test_endpoint.go)

<br />

Most outside resources used are listed in [external_resources.txt](external_resources.txt)

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
