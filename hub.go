package main

import "log"
import "time"
import "github.com/gorilla/websocket"
import "net/http"

import "encoding/json"

import "crypto/md5"
import "fmt"

import "strconv"
import "math/rand"
import "github.com/satori/go.uuid"


const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
        MAX_MSG_SIZE = 5000
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 1024

        HISTORY_CAP = 200
)


type UniqueRand struct {
    generated map[int]bool
}

func (u *UniqueRand) Int() int {
    for {
        i := rand.Int()
        if !u.generated[i] {
            u.generated[i] = true
            return i
        }
    }
}

type Message struct {
        *Event
        Id    string  
        data      []byte
	Content   string  `json:"content"`
	Room  string `json:"room"`
        Color    string `json:"Color"` 
       	Author    string `json:"author"`
        Timestamp int64   `json:"timestamp"`   
        
        
}
type Author struct {
        ID       int  `json:"id"`
	Username string `json:"username"`
        Color    string `json:"Color"`
	*Room
}

type Auteur struct {
	Username string `json:"username"`
	
}


type Subscription struct {
	Conn  *Connection
        //*Team
	*Room
	//Email string `json:"Email"`
        
}

type Connection struct { /////////////////////// Client
	// The websocket connection.
	Hub *Hub
	Id  string
	ws  *websocket.Conn
        *Room
        //Session  *Session
        //UserId   string
	// Buffered channel of outbound messages.
	send chan []byte
        Out  chan []byte
	
	Events map[string]EventHandler
}



// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	rooms map[string]map[*Connection]bool
        roomis map[string]*managedRoom
        Teamis map[string]map[string]*managedTeam
	emails map[string]string
        Sockets  map[*Connection]*Room
        Sackets  map[*Connection]*Author
        tomis  map[string]map[string]*managedTeam
        *Session
	//emails map[string]map[*connection]bool
	// Inbound messages from the connections.
	//broadcast chan *Message 
        Channels  map[string]Channel 
        History     []*Event
        HistorySess []*Session
        Teams       []*Team
        Emojis      []*Emoji
        Sessions    map[string]*Session
	//broadcast chan string
	// Register requests from the connections.
	Clients   []Connection
        Subscriptions []*Subscription
        raamis   map[string]*managedRoom
	//Broadcast chan []byte
	nextID    int
	Users     map[string]*Author
        //Join      chan *roomReq 
        public    chan *raamReq
        Posts     []*Post
        Members   []*TeamMember
        TeamMembers []*Team
        TempHistory map[int64]map[int64]bool
        clients   map[*Connection]bool
	register chan Subscription
        users       map[string]*Author
	// Unregister requests from connections.
	unregister chan Subscription
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Avatar(email string, size uint) string {
	hash := md5.Sum([]byte(email))
	return fmt.Sprintf("https://www.gravatar.com/avatar/%x?s=%d", hash, size)
}

func autoId() (string) {

	return uuid.Must(uuid.NewV4()).String()
}



func newHub() *Hub {
	return &Hub{
		
         	register:   make(chan Subscription),
                roomis:     make(map[string]*managedRoom),
                tomis:      make(map[string]map[string]*managedTeam),
                Teamis:     make(map[string]map[string]*managedTeam),
                Sackets:    make(map[*Connection]*Author),
                raamis:     make(map[string]*managedRoom), 
                Channels:   make(map[string]Channel), 
                Sockets:    make(map[*Connection]*Room),
                //Join:       make(chan *roomReq),
                public:     make(chan *raamReq), 
                Sessions:    make(map[string]*Session),
                History:    make([]*Event, 0),
                Posts:      make([]*Post, 0),
                Teams:      make([]*Team, 0),
                Emojis:     make([]*Emoji, 0),
                HistorySess:    make([]*Session, 0),
		//Broadcast:  make(chan []byte, 256),
		Clients:    make([]Connection, 0),
                Members:     make([]*TeamMember, 0),
                TeamMembers: make([]*Team, 0),
                TempHistory: make(map[int64]map[int64]bool),
                Subscriptions: make([]*Subscription, 0),
		Users:      make(map[string]*Author),
                users:      make(map[string]*Author),  
		unregister: make(chan Subscription),
		rooms:      make(map[string]map[*Connection]bool),
                emails:     make(map[string]string),
                clients:    make(map[*Connection]bool), 
	}
}

func (h *Hub) AddClient(c Connection) *Hub {
	h.Clients = append(h.Clients, c)
	log.Println("add new client ", c.Id)

	return h

}

func (h *Hub) AddFirstRoomis() map[string]*managedRoom {
	h.roomis["@me"] = &managedRoom{
				room:  &Room{Name: "@me"},
                                        count: 1,
					 
			}
        h.roomis["606205760728465428"] = &managedRoom{
				room:  &Room{Name: "606205760728465428"},
                                        count: 2,
		 	 
			} 


	return h.roomis

}

func (h *Hub) GetTeamis(tims []*Room, tams []*Team) map[string][]*Room {
         
        //y := make(map[string][]*Team)
        x := make(map[string][]*Room)
       // tom map[string]*managedTeamCount
        y := make(map[string][]*Team)

        for i := range tams {
        //fmt.Println(tams[i], tims[i])

        x[tams[i].Displayname] = append(x[tams[i].Displayname], tims[i])
        y[tams[i].Displayname] = append(y[tams[i].Displayname], tams[i])
        //y[tams[i].Displayname] = append(y[tims[i].Displayname], tims[i])
        }
        toms := &managedTeamCount{count: 1 , rooms: x, teams: y}
        //mi := &managedTeamCount{rooms: paii, teams: poiu}
	

                //log.Println(len(tims))
	

        return toms.rooms
	

}


func (h *Hub) GetUrls(tims []*Room, tams []*Team) map[string][]interface{} {
         
        //y := make(map[string][]*Team)
        //x := make(map[string][]*Room)
       // tom map[string]*managedTeamCount
        ud := make(map[string][]interface{})

        //yi := make(map[string][]interface{})

        for i := range tams {
        
        //ud = map[string]interface{}{"url": tams[i].Displayname + tims[i].Displayname }
        ud[tams[i].Displayname] = append(ud[tams[i].Displayname], "/" + tams[i].Displayname + "/channels/" + tims[i].Displayname + "/") 
        //x[tams[i].Displayname] = append(x[tams[i].Displayname], tims[i])
        //y[tams[i].Displayname] = append(y[tams[i].Displayname], tams[i])
        //y[tams[i].Displayname] = append(y[tims[i].Displayname], tims[i])
        }
        //toms := &managedTeamCount{count: 1 , rooms: x, teams: y}
        //mi := &managedTeamCount{rooms: paii, teams: poiu}
	

                //log.Println(len(tims))
	

        return ud
	

}

func (h *Hub) GetEmojis(emojis []*Emoji) map[string]*Emoji {

        //y := make(map[string][]*Team)
        //x := make(map[string][]*Room)
       // tom map[string]*managedTeamCount
        ud := make(map[string]*Emoji)

        //yi := make(map[string][]interface{})

        for _, emo := range emojis {
        
        //ud = map[string]interface{}{"url": tams[i].Displayname + tims[i].Displayname }
        ud[emo.Keyi] = emo 
        //x[tams[i].Displayname] = append(x[tams[i].Displayname], tims[i])
        //y[tams[i].Displayname] = append(y[tams[i].Displayname], tams[i])
        //y[tams[i].Displayname] = append(y[tims[i].Displayname], tims[i])
        }
        //toms := &managedTeamCount{count: 1 , rooms: x, teams: y}
        //mi := &managedTeamCount{rooms: paii, teams: poiu}
	

                //log.Println(len(tims))
	

        return ud
	

}










 
func (h *Hub) updateRoomis(name string)  map[string]*managedRoom {
	
       
	 if _, ok := h.roomis[name]; !ok {
             test := make(map[string]*managedRoom) 
             log.Println(test)
             h.roomis[name] = &managedRoom{room:  &Room{Name: name}}        
					 
			
         }	
	
		
	return h.roomis
	

}




func NewWebSocket(h *Hub, w http.ResponseWriter, r *http.Request) (*Connection, error) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("[ERROR | SOCKET CONNECT] %v", err)
		return nil, err
	}
	// conn.SetWriteDeadline(time.Now().Add(MSG_TIMEOUT))
	co := &Connection{
		Hub:   h,
		ws:   ws,
		send: make(chan []byte),
                Out:  make(chan []byte),
		Events: make(map[string]EventHandler),
	}
	go co.Reader()
	go co.Writer()
	return co, nil
}

func (co *Connection) Reader() {
	defer func() {
		co.Hub.Unregister(co)
		co.ws.Close()
	}()
	co.ws.SetReadLimit(MAX_MSG_SIZE)
	for {
		_, message, err := co.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("[ERROR] %v", err)
			}
			break
		}
		event, err := NewEventFromRaw(message)
		if err != nil {
			log.Printf("[ERROR | MSG] %v", err)
		} else {
			log.Printf("[MSG] %v", event)
		}
		if action, ok := co.Events[event.Name]; ok {
			action(event)
		}
	}
}


func (co *Connection) Writer() {
	for {
		select {
		case message, ok := <-co.send:
			if !ok {
				co.ws.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := co.ws.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			w.Close()
		}
	}
}


func (h *Hub) Unregister(co *Connection, conerr ...bool) {
	log.Printf("[SOCKET DISCONNECTED]")
	if action, ok := co.Events["disconnected"]; ok && len(conerr) == 0 && h.Sockets[co] != nil && h.Sackets[co] != nil {
		action(&Event{
			Name: "disconnected",
			Data: map[string]interface{}{
				"name":     h.Sockets[co].Room.Author.Username,
				"nclients": len(h.Sockets),
			},
		})
	}


        
	delete(h.Sockets, co)
        delete(h.Sackets, co) 
        
	co.ws.Close()
}

func (h *Hub) Register(room *Room, co *Connection) {

        if co.Room == nil{
            co.Room = room

        }

        log.Printf("[SOCKET CONNECTED]")

        //h.Sockets[s.Conn] = s.Room
        connections := h.rooms[room.Displayname]

			if connections == nil {
				connections = make(map[*Connection]bool)
				h.rooms[room.Displayname] = connections
			}

			h.rooms[room.Displayname][co] = true
        h.Sackets[co] = room.Author

}








func (h *Hub) Broadcast(message []byte, room *Room) {
        log.Println("sackets", h.Sockets)
connections := h.rooms[room.Displayname]
        

        //log.Println("co", connections)
       
        //log.Println("so", c.Sockets)
        log.Println("sv", connections)
        
	for s := range connections {
                
               
                log.Printf("s %v", s)
		select {
		case s.send <- message:
		default:
			h.Unregister(s, true)
		}
	}


}





func (c *Connection) write(mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, payload)
}






func NewMessageFromRaw(rawData []byte) (*Message, error) {
	mess := &Message{}
	err := json.Unmarshal(rawData, mess)
	return mess, err
}

func (m *Message) ToJson() string {
	b, err := json.Marshal(m)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}


func (m *Message) Raw() []byte {
	raw, _ := json.Marshal(m)
	return raw
}

// readPump pumps messages from the websocket connection to the hub.
// readPump pumps messages from the websocket connection to the hub.
func (s Subscription) readPump() {
	c := s.Conn
        

	defer func() {
		c.Hub.unregister <- s
		c.ws.Close()
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, msg, err := c.ws.ReadMessage()
               
      		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("[ERROR] %v", err)
			}
			break
		}
		event, err := NewEventFromRaw(msg)
                
                             
		log.Printf("ddddd %v", event)
		if err != nil {
			log.Printf("[ERROR | MSG] %v", err)
		} else {
			log.Printf("[MSG] %v", event)
                        
		}
		if action, ok := c.Events[event.Name]; ok {
                        log.Printf("event %v", event.Name) 
			action(event)
		}
               

               
                
	}
}

func (s *Subscription) writePump() {
	c := s.Conn
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
                        if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func (c *Connection) SetHandler(event string, action EventHandler) *Connection {
	c.Events[event] = action
	return c
}


func (h *Hub) AppendSess(session *Session) {
        //log.Println("evi", event)
	
	h.HistorySess = append(h.HistorySess, session)
}




func (h *Hub) FetchSess(ident func(*Session) bool) (*Session, string) {
        
	for i, msg := range h.Sessions {
        //log.Println("evi", msg)
		if ident(msg) {
			return msg, i
		}
	}
	return nil, ""
}

func (h *Hub) GetSessByID(id int64) (*Session, string) {
	return h.FetchSess(func(e *Session) bool {
               
                //n := strconv.FormatInt(e.ID, 10)
               

		return (e.ID == id)
	})
}

func (h *Hub) GetSessByAuthor(author string) (*Session, string) {
	return h.FetchSess(func(e *Session) bool {
               
                //n := strconv.FormatInt(e.ID, 10)
               

		return (e.Authori == author)
	})
}

func (h *Hub) groupByPosts(maps []map[string]*Post, key string) map[string][]map[string]*Post {
  groups := make(map[string][]map[string]*Post)
  for _, m := range maps {
    k := m[string(key)] // XXX: will panic if m[key] is not a string.
    groups[k.Rootid] = append(groups[k.Rootid], m)
  }
  return groups
}



func (h *Hub) groupBy(maps []map[string]interface{}, key string) map[string][]map[string]interface{} {
  groups := make(map[string][]map[string]interface{})
  for _, m := range maps {
    k := m[key].(string) // XXX: will panic if m[key] is not a string.
    groups[k] = append(groups[k], m)
  }
  return groups
}


func (h *Hub) AppendHistory(event *Event) {
	if len(h.History) > HISTORY_CAP {
		h.History = append(h.History[len(h.History)-HISTORY_CAP:], event)
		return
	}
	h.History = append(h.History, event)
}



func (h *Hub) FetchPost(ident func(*Post) bool) (*Post, int) {
        
	for i, msg := range h.Posts {
         
		if ident(msg) {
			return msg, i
		}
	}
	return nil, -1
}

func (h *Hub) GetPostByID(id int64) (*Post, int) {
	return h.FetchPost(func(e *Post) bool {
               
                n, err := strconv.ParseInt(e.Id, 10, 64)
               if err == nil {
                  fmt.Printf("%d of type %T", n, n)
                  }

		return (n == id)
	})
}

func (h *Hub) DeletePostByID(id int64) *Post {
	msg, i := h.GetPostByID(id)
	if msg == nil {
		return nil
	}
	h.Posts = append(h.Posts[:i], h.Posts[i+1:]...)
	return msg
}

func (h *Hub) FetchMessage(ident func(*Event) bool) (*Event, int) {
        log.Println("evi", h.History)
	for i, msg := range h.History {
         
		if ident(msg) {
			return msg, i
		}
	}
	return nil, -1
}

func (h *Hub) GetMessageByID(id int64) (*Event, int) {
	return h.FetchMessage(func(e *Event) bool {
               
                n, err := strconv.ParseInt(e.Message.Id, 10, 64)
               if err == nil {
                  fmt.Printf("%d of type %T", n, n)
                  }

		return (n == id)
	})
}







func (h *Hub) DeleteMessageByID(id int64) *Event {
	msg, i := h.GetMessageByID(id)
	if msg == nil {
		return nil
	}
	h.History = append(h.History[:i], h.History[i+1:]...)
	return msg
}

func (h *Hub) EnqueueTempHistory(id int64) {
	now := time.Now().UnixNano()
	if len(h.TempHistory[id]) == 0 {
		h.TempHistory[id] = map[int64]bool{
			now: true,
		}
	} else {
		h.TempHistory[id][now] = true
	}
	time.AfterFunc(10*time.Second, func() {
		delete(h.TempHistory[id], now)
	})
}

func (h *Hub) TempHistoryLength(id int64) int {
	if h.TempHistory[id] == nil {
		return 0
	}
	return len(h.TempHistory[id])
}


func (h *Hub) Loginnn(co *Connection, author *Author, addr string) {
        ou := &UniqueRand{} 
        Connes := ou.generated
        if Connes == nil {
          Connes = make(map[int]bool)
	        ou.generated = Connes
        }
        ou.generated[125] = true
        id := int64(ou.Int())
        hash := CreateHash(strconv.FormatInt(id, 10), author.Username, addr)
        session := &Session{ID: id, Authori: author.Username, Hash: hash}  
	
	log.Println("create cookie")
	go func() {
		co.send <- (&Event{
			Name: "createCookie",
			Data: (&http.Cookie{
				Name:     "cookie-name",
				Value:    hash,
				Expires:  time.Now().Add(SESSION_TIMEOUT),
				HttpOnly: false,
			}).String(),
		}).Raw()
	}()
        h.Session = session
        store.CreateSession(session)
	h.Sessions[session.Hash] = session
	time.AfterFunc(SESSION_TIMEOUT, func() {
		delete(h.Sessions, session.Hash)
	})
}




       







         

	
	


