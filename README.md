Microservices

1. Monolithic applications
2. Distibuted application 
3. Mircroservices

Do one thing and do it well - the unix philosophy

1. Monoliths

Everytibng is handled by one applicaiton
1. user authntication, sending email, logging, all business logic

2. Microservices
1. Breaking monolith up from funcitons/ package to completely separate programs
2. communicationvis json/rest, rpc, frpc, and vocer a mesageing queue
3. easire to scale
4. Easier to maintain
5. harder to write

-----------------
A front end web applicatoin that connect to 5 microservis 
1. Vroker- optional single ponts of enty to mirco servis 
2. authenticaton: postagres
3. Mail - sends emai  with a specific template
4. Listener - consumer message in rabbitmq and initate a process

we will commuicaate from the bwtween microsecices using 
rest api with json as traspost
sending and receiving using rpc 
sending and receiving using grpc

initiatin and responding to events using advcanced message queing protocaol(amqp)