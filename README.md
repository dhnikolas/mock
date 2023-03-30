### Self-hosted HTTP Mock Service 
Run with docker:

    docker run -v `pwd`:/var/mock/ -p 8111:8111 dhnikolas/mock:v2.2

Usage with application UI:

    http://localhost:8111/ui/

Configure endpoints:

<img width="1534" alt="Снимок экрана 2023-03-30 в 02 55 50" src="https://user-images.githubusercontent.com/15860590/228693972-36d7483d-ecf5-4f21-a894-0aeb2ad5f9bf.png">

Make request and get response: 

    $ curl http:/localhost:8111/api/v1/pam
    {"project_id","s98fsufhs9fusnf9si"}
    

Show request logs: 

<img width="1688" alt="Снимок экрана 2023-03-30 в 03 01 00" src="https://user-images.githubusercontent.com/15860590/228694483-812d90a8-76e1-4a97-b1ec-3803b96107f1.png">