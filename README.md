# hackathon-may

## Obecny backend
Dokumentacja routingu znajduje się w autogenerowanym pliku [routes.md](docs/routes.md) (`make documentation`).
Aplikacja wspiera bardzo prosty CRUD - tworzenie/update i czytanie metadanych artykułów oraz administracyjny delete.

W kodzie zaszyte są również mock calle do zewnętrznych serwisów, które mają spełniać założenia klienckie:
- auth dla administratorów
- propagacja notyfikacji - generowanie notyfikacji dla użytkowników z listy mailingowej, gdy zostaje stworzony nowy artykuł
- dołączanie `reklam` do nowo stworzonego artykułu

## Flow charts

_Cała komunikacja między monolitem a serwisami nie musi być synchroniczna/dwukierunkowa/stanowa etc..._
_Connection flowy również nie są określone na sztywno, możemy omówić jak połączyć te mikroserwisy, żeby zapewnić większą współpracę między wami_
_Dorzucajcie też swoje pomysły na fukcjonalności, żeby rozszerzyć koncept_

### **User Notification Service**
```mermaid
sequenceDiagram
    participant User
    participant Monolith
    participant UserNotificationSvc
    participant UserGroup

    User->>Monolith: PUT /article (create) -> articleID
    Monolith->>UserNotificationSvc: Generate new notification about <br/> article with articleID
    UserNotificationSvc->>UserGroup: Notification  <br/> preferably via UDP-based protocol
    UserNotificationSvc->>Monolith: Respond if errors, retry/add fallback
    Monolith->>User: 200 OK
```

### **Authentication Service**
```mermaid
sequenceDiagram
    participant User
    participant Monolith
    participant AuthSvc

    User->>Monolith: DELETE /article/{articleID}
    Monolith->>AuthSvc: Check if user has admin privileges
    AuthSvc->>Monolith: Respond with privilege boolean
    Monolith->>User: Allow/Disallow operation
```

### **Comercials Services**
```mermaid
sequenceDiagram
    participant User
    participant Monolith
    participant ComercialGeneratorSvc
    participant ComercialAttacherSvc

    User->>Monolith: PUT /article (create) -> articleID
    Monolith->>ComercialGeneratorSvc: Generate commercial for article <br/> with articleID -> (articleID, comercialID)
    ComercialGeneratorSvc->>ComercialAttacherSvc: Merge article with articleID with comercial
    ComercialAttacherSvc->>Monolith: Respond with optional(err)
    Monolith->>User: 200 OK

```
