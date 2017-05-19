# ICA07
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

