Kommentar: 
- Alle i gruppen deltok på ICAen, men Magnus Erdvik gjorde mest på ICA07 mens resten av gruppen fokuserte på 2. utkast av de andre ICAene. Vi mente dette var den beste måten å fordele oppgavene på for å bli ferdig i tide. 

Deltakere:

Erdvik, Magnus

Lie, Eva Kristine

Nguyen, Philip

Tellefsen, Erlend Frøysnes

Van Dijk, Richard

Younas, Osman


                                                   ICA07 - Gruppe 11
Oppgave 1.

a) Lag en UDP klient og en UDP tjener
- Se “udpclient.go” og “udpserver.go”. Kjøres ved å ha to terminaler der den ene er server og den andre er klient. Vi lagde en windows executables av programmene; udpclient.exe og udpserver.exe ("./udpclient" og "./udpserver" i terminalen). Eventuelt kan programmene kjøres med kommando "go run udpserver.go kryptering.go" og "go run udpclient.go kryptering.go". Implementasjon av kryptering er gjort i en senere oppgave og må dermed være med for å kjøre prgrogrammene. 

b) Send over “hemmelig” melding “Møte Fr 5.5 14:45 Flåklypa”. 
- klient: http://imgur.com/Ea9Xiha
- server: http://imgur.com/4v8H8Nl

c) Studer kommunikasjon i Wireshark

i) på det lokale grensesnittet (obs! windows brukere kan eventuelt droppe det)
 1) Hvor mange prosent av data, som blir sendt over en nettverksforbindlese, er protokoll-data, dvs. data, som er nødvendig for å transportere meldingen fra bruker over nettverksforbindelsen? 
 - Ved bruke RawCap kunne vi se loopback-interfacet på windows maskinen. Ved å sende den hemmelige meldingen: “Møte Fr 5.5 14:45 Flåklypa” over lokal node fikk vi dette opp i wireshark: http://imgur.com/hm45Md3
 
 - Som vist på toppen av bildet er totalt 57 bytes fanget opp. Av dem kan vi se at selve meldingen er 29 bytes, dermed må protocol data være de resterende 28 bytes. Protocol data utgjør derfor ca. 49% av datamengden (28/57). 
 
 2) Hvor stor kan en UDP pakke være? Begrunn. 
 - Ifølge wikipedia (https://en.wikipedia.org/wiki/User_Datagram_Protocol under “Packet structure”, “Length”) er den teoretiske størrelsen på en udp pakke 65 535 bytes, men den faktiske størrelsen er på 65 507 bytes. Dette kommer av at IPv4 trenger noen av bytene til headeren (8 bytes til UDP header og 20 bytes til IP header). 
 
- Over nettverket blir størrelsen til UDP pakken holdt tilbake av Ethernet 2 sin MTU (maximum transmission unit). MTU er på maks 1500 bytes, noe som gjør at pakker over 1500 bytes blir fragmentert i flere pakker, før den blir “satt sammen” igjen ved destinasjonen. 

- Ved å ta vekk headerdata som nevnt over, blir den største udp pakken man kan sende over nettet 1472 bytes (1500 - (8 + 20)). MAC headeren (14 bytes) er ikke en del av Ethernet 2 sin MTU.

ii) over NIC
1) Hvor mange prosent av data, som blir sendt over en nettverksforbindlese, er protokoll-data, dvs. data, som er nødvendig for å transportere meldingen fra bruker over nettverksforbindelsen?
- Her er et eksempel på en udp pakke som er sendt over wifi nettverket: http://imgur.com/8h6COmF

- Her er det 42 bytes med protocol data, ettersom det totalt blir fanget 126 bytes og 84 av dem er ikke protocol data. Protocol dataen er større i dette eksempelet ettersom dataen blir transportert over wifi og det dermed trengs mer data for å finne fram til destinasjonen. 

2) Hvilken filter må du bruke for å filtrere ut relevante meldinger?
- Ved å bruke feltet som vist på bildet (http://imgur.com/IxgmnO2) kan man i wireshark filtrere ut hvilke typer pakker man ønsker skal bli vist. For eksempel ønsket vi kun se udp pakker og skrev dermed “udp” i feltet.  Grunnen til at quick og mdns protokollene vises er siden de er en type udp protocol. 

- For å kun se tcp pakker kan man skrive inn "tcp" i feltet. Liknende filtre man kan bruke er “ip” som kun viser Ipv4, “ip.addr == <ønsket ip adresse” for å kun se pakker sendt fra den ip adressen, osv. Filter i wireshark er et veldig bra redskap for å kunne filtrere ut kun relevante data. 

 3) Hva er forskjell mellom data som ble sendt over NIC (ditt nettverkskort, mest sannsynlig trådløst) og de som ble sendt innenfor din lokale node? Illustrer gjerne med “screenshots” eller log-filer. Forklar.
 - Over lokal node: http://imgur.com/lwKxxgQ
 - Over NIC: http://imgur.com/a4Evwk8
 
 - Over lokal node kan vi se at det ikke er et felt for “Ethernet 2” slik det er på pakken sendt over wifi. Det er fordi pakken som er sendt over lokal node er raw data og har ikke behov for noe data for  å transporteres over datalink laget i OSI-modellen. 
 
Oppgave 2. 

a) Lag en TCP klient og en TCP tjener 
- Se “tcpclient.go” og “tcpserver.go”. Kjøres ved å først starte tcp server med “go run tcpserver.go” en terminal og så “go run tcpclient.go” i en annen terminal. (Anbefales å stoppe server (Ctr + c) først ellers vil det printes ut “Melding:” i en loop). 

b) Studer i Wireshark (vinduer og andre parametre) 

i) Hva er forskjellig fra UDP? 
- Ved å sende den samme meldingen (Møte Fr 5.5 14:45 Flåklypa) er det blitt sendt 5 pakker over lokal node istedenfor bare 1 som ved bruk av UDP protokollen: http://imgur.com/w6mW8PI 

- Dette kommer av TCP sin “three-way handshake” (https://www.youtube.com/watch?v=Vdc8TCESIg8): I pakke 1 vil klienten spørre serveren om den ønsker å gjøre en tilkobling [SYN]. I pakke 2 sender serveren en godkjenning av tilkoblingen til klienten [SYN, ACK]. I pakke 3 sender klienten en godkjenning til serveren om at tilkoblingen er gjort [ACK]. I pakke 4 sender klienten selve meldingen [PSH] sammen med forespørsel om at all dataen har kommet frem [ACK]. I pakke 5 sender serveren godkjenning om at all dataen har kommet frem [ACK]. 

- Dette viser en distiktiv forskjell mellom UDP og TCP ettersom UDP kun sender pakkene til destinasjonen uten “three-way handshake” som TCP gjør. Dette gjør det mulig for pakker sendt over TCP å bli sendt på nytt om den ikke når fram ettersom det er en tilkobling mellom klient og server, noe som ikke skjer med UDP. Denne tilkoblingen er det som gjør at TCP er så pålitelig, men er også grunnen til at den bruker mye mer data (ca. 60%) og er tregere enn UDP. 

ii) Hvor stor kan en TCP pakke være?
- Hvor stor en TCP pakke kan bli er bestemt av MSS (Maximum Segment Size) i headeren. Dette blir satt under nevnte "three-way handshake". I eksempelet over er den satt på 65495 bytes (under options, MSS Value): http://imgur.com/jnVbV6T

-  Men i likhet til UDP er Størrelsen limitert av Ethernet 2 sin MTU som er på 1500 bytes. Pakker større enn MTU blir da fragmentert.

iii) Hva er fragmentering, hvorfor oppstår det og hvordan håndterer man det?
- Fragmentering oppstår av at datamengen som skal bli transportert er for stor til å blir fraktet i en pakke, og dermed blir delt opp i flere pakker. TCP skiller seg fra UDP når det gjelder fragmentering ettersom TCP vil holde styr på at all dataen som skal bli sendt har kommet fram. Om data ikke er kommet fram vil TCP passe på at klienten sender dataene på nytt helt til alle dataene er kommet fram. TCP vil også passe på at de fragmenterte datapakkene blir satt sammen i riktig rekkefølge, noe UDP heller ikke gjør. Dermed håndteres fragmenteringen direkte av TCP. UDP tar istedenfor i bruk checksum og flags for å merke og eventuelt forkaste pakker som ikke er fullkommne.

iv) I hvilke brukerscenarier bruker man UDP og i hvilke TCP?
- TCP brukes når man må være sikker på at pakkene som blir sendt over kommer fram til mottaker og at pakkene kommer i riktig rekkefølge. For eksempel webapplikasjoner, epost og chattesystemer.

- UDP brukes når det ikke er så viktig at dataen alltid kommer fram og i riktig rekkefølge, men at det skjer raskt. For eksempel videostreaming, skype, ulike “fast pace” spill; generelt der det ikke er behov for påliteligheten til tcp og høy fart er ønsket. 

Oppgave 3.

a) Implementer kryptering for eksemplet fra Oppgavene 1 og 2 
- Se “kryptering.go”; implementert i udp server/client. (Terminal 1: "go run udpclient.go kryptering.go", Terminal 2: "go run udpserver.go kryptering.go"). 








 

