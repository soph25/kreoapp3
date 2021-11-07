package main

import (
    "fmt"
    "net/http"
    "strings"
    "time"
    "log"
    "os"
    "reflect" 
    "regexp"
    "strconv"
    "github.com/gorilla/sessions"
    "html/template"  
    "database/sql"
    _ "github.com/lib/pq" 
    "github.com/gorilla/mux"
    "encoding/json"
    "io"  
    "golang.org/x/crypto/bcrypt"
    "github.com/russross/blackfriday"
    "github.com/grokify/html-strip-tags-go"
)

type Group struct {
	Key   interface{}
	Group []interface{}
}


type Server struct {
	Store           Store
	
	Router          *mux.Router
	
}

type Routes struct {
	Root    *mux.Router
        PostsForChannel *mux.Router // 'api/v4/channels/{channel_id:[A-Za-z0-9]+}/posts'
}

type API struct {
	
	BaseRoutes          *Routes
}

type Emoji struct {
        Id         string `json:"id"` 
	Keyi       string `json:"keyi"`
	Valeur     string `json:"valeur"`
	Descriptor string `json:"descriptor"`
}


type Post struct {
	Id            string          `json:"id"`
	Createat      int64           `json:"createat"`
        Userid        string          `json:"userid"`
	Channelid     string          `json:"channelid"`
        ReplyCount    int64           `json:"reply_count" db:"-"`
	Rootid        string          `json:"rootid"`
        Parentid      string          `json:"parentid"`
	Message       string          `json:"message"`
	Color         string          `json:"color"`
        Typo          string          `json:"typo"`
        *Props        
}

type PostPatch struct {
	Message      *string          `json:"message"`
	Props        *StringInterface `json:"props"`
	
	
}


type Team struct {
	Id              string `json:"Id"`
	Name            string `json:"Name"`
	Description     string `json:"Description"`
	Email           string `json:"Email"`
        Displayname     string `json:"Displayname"`
	CreateAt        int64  `json:"CreateAt"`
	namee           string `json:"namee"`
}

type ChannelList []*Room

type TeamMember struct {
	Teamid   string `json:"team_id"`
	Userid   string `json:"user_id"`
	Roles    string `json:"roles"`
	DeleteAt int64  `json:"delete_at"`
}


type Room struct {
        Id   string `json:"Id"`
	Name  string `json:"Name"`
        CreateAt      int64  `json:"CreateAt"`
        Teamid        string `json:"Teamid"`
	Type          string `json:"type"`
	Displayname   string `json:"Displayname"`
        *Author
        *Session
        Authorid      string `json:"authorid"`
        //connections    map[*Connection]bool
}
type Channel []Room

type TeamList []*Team


type raamReq struct {
	// Name of the lobby the request goes to.
	namee string

                
        room Room
	// Reference to the connection, which requested.
	conn Connection
}

type Session struct {
	ID     int64   `json:"id"`
	Authori string `json:"authori"`
	Hash   string  `json:"hash"`
        
}


type User struct {
    Id             int 
    Email          string
    CreatedAt      time.Time
    UpdatedAt     time.Time 
    HashedPassword []byte
    Password       string
}

type roomReq struct {
	// Name of the lobby the request goes to.
	name string

                
        room *Room
	// Reference to the connection, which requested.
	conn *Connection
}

type managedTeam struct {
	// Reference to room.
        
	teams map[string][]*Team
        rooms map[string][]*Room
	// Member-count to allow removing of empty lobbies.
	
}

type managedTeamCount struct {
	// Reference to room.
        count int
	teams map[string][]*Team
        rooms map[string][]*Room
	// Member-count to allow removing of empty lobbies.
	
}

type managedSession struct {
	// Reference to room.
        
	room *Room
        team *Team
        session *Session
	// Member-count to allow removing of empty lobbies.
	
}


type managedRoom struct {
	// Reference to room.
	room *Room

        
	// Member-count to allow removing of empty lobbies.
	count uint
}

var appSession *Session 

var appTeam *Team 

var Srv *Server

var storaa Stora 



var (
    // key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
    key = []byte("super-secret-key")
    stori = sessions.NewCookieStore(key)
    channelas = map[string]*Channel{}
    
)

func Init(root *mux.Router) *API {
      api := &API{
		
		BaseRoutes:          &Routes{},
	}
api.BaseRoutes.Root = root

root.Handle("/api/v4/{anything:.*}", http.HandlerFunc(api.Handle404()))

	return api



}
func strip_tags(content string) string {
	re := regexp.MustCompile(`<(.|\n)*?>`)
	return re.ReplaceAllString(content,"")
}



func escaped(b byte) int {
	return strings.IndexByte("\\!\"#$%&'()*+,./:;<=>?@[]^_`{|}~-", b) 
}

func testEsc(s string) string {
     bi := []byte(s) 
     y := make([]byte, 0)
     
     for _, b := range bi {
           if escaped(b) > -1 {
              
             //return bi
           } else {
             y = append(y, b)
             //return y
           }
     }
     return string(y) 

}


func (api *API) Handle404() func(http.ResponseWriter, *http.Request) {
   return func(w http.ResponseWriter, r *http.Request) {
       w.WriteHeader(http.StatusNotFound)
       log.Println("custom 404")
   }
 
}

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}



func getField(v *Session, field string) string {
    r := reflect.ValueOf(v)
    f := reflect.Indirect(r).FieldByName(field)
    return string(f.String())
}

func GeneratePasswordHash(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}

func sum(so []Room, c chan []Room, r Room) {
	
	
	so = append(so, r )
	
	c <- so // send sum to c
}

func TeamMapToJson(u map[string]*Team) string {
	b, err := json.Marshal(u)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}
func JsonStory(r io.Reader) ([]Team, error) {
	d := json.NewDecoder(r)
        var ms = []Team{}
        
	
        if err := d.Decode(&ms); err == io.EOF {
            log.Fatal(err)
        } else if err != nil {
            log.Fatal(err)
        }
	
	return ms, nil
}
func TeamListToJson(t []*Team) string {
	b, err := json.Marshal(t)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func Index(vs []*Team, t string) string {
    for _, v := range vs {
        if v.Displayname == t {
            return v.Displayname
        }
    }
    return ""
}

func Indexo(vs []*Room, t string) string {
    for _, v := range vs {
        if v.Displayname == t {
            return v.Displayname
        }
    }
    return ""
}

func Indexou(vs []*Room, t string) *Room {
    for _, v := range vs {
        if v.Displayname == t {
            return v
        }
    }
    return nil
}

func Filter(vs []*TeamMember, tamid string) *TeamMember {
    
    for _, v := range vs {
        if v.Teamid == tamid {
            return v
        }
    }
    return nil
}

func Find(a []*TeamMember, x string) *TeamMember {
        for _, n := range a {
                if x == n.Teamid {
                        return n
                }
        }
        return nil
}

func Contains(a []*TeamMember, x string) bool {
    for _, n := range a {
        if x == n.Teamid {
            return true
        }
    }
    return false
}

func (o *PostPatch) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	}

	return string(b)
}

func PostPatchFromJson(data io.Reader) *PostPatch {
	decoder := json.NewDecoder(data)
	var post PostPatch
	err := decoder.Decode(&post)
	if err != nil {
		return nil
	}

	return &post
}

func FetchListTeam([]*Team) string {
        tams, _ := store.GetTeams()
        for i := 0; i < len(tams); i++ {
		return tams[i].Displayname
	}

	return ""
}

func (h *Hub) FetchTeam(ident func(*Team) bool) (*Team, int) {
        tams, _ := store.GetTeams()
	for i, msg := range tams {
        log.Println("evi", msg)
		if ident(msg) {
			return msg, i
		}
	}
	return nil, -1
}
func (o *Team) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func (o *Post) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func PostListToJson(t []*Post) string {
	b, err := json.Marshal(t)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}


func (emoji *Emoji) ToJson() string {
	b, _ := json.Marshal(emoji)
	return string(b)
}

func ArrayToJson(objmap []*Post) string {
	b, _ := json.Marshal(objmap)
	return string(b)
}


func (h *Hub) Secret() func(http.ResponseWriter, *http.Request) {
       
	return func(w http.ResponseWriter, r *http.Request) {
                     segs := strings.Split(r.URL.Path, "/")
    //log.Println(segs[3], "The cake is a lie!")  
    room := segs[3] 
    session, _ := stori.Get(r, "cookie-name")
    
   
    //h.updateRoomis(vars["room"])
    
        
    if cookie, err := r.Cookie("cookie-name"); err == http.ErrNoCookie {
                log.Println(cookie)
                w.Header()["Location"] = []string{"/login"}
		w.WriteHeader(http.StatusTemporaryRedirect)

    } else if auth, ok := session.Values["species"].(bool); !ok || !auth {
    // Check if user is authenticated
    
        log.Println(session.Values["species"].(bool))  
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    } else if session.Values["username"] != nil {
    //log.Println(session.Values["species"].(bool))
    user := session.Values["username"] 
    
        log.Println(user)

         datass := map[string]interface{}{
		"Host": r.Host,
    } 

    if datass != nil {
    datass["UserData"] = session.Values["username"]
    
    numi := datass["UserData"].(string)
    //log.Println("aaa", h.chackManager(room).room.Name) 
    vars := mux.Vars(r)
    //log.Println("room", vars["room"])
    //log.Println("team", vars["team"]) 
    uff, _ := store.GetTeam("pppppppp")
    n, present := h.roomis[room]
    log.Println("neeeeeeeeeeeeeeeeeeeeee", n)
        if present == false {
             http.Redirect(w, r, "/" + uff.Displayname + "/channels/@me/", http.StatusFound)
        }  

	//tams, err := store.GetTeams()
       //log.Println(err)
    pit, errt := store.GetTeamByEmail(numi)
        if errt != nil {
        log.Println("oooo", errt) 
        } 
        if pit != nil {
        //log.Println("handleroooopitttttttttttttttt", pit.Displayname)  
     
        
        
        if vars["team"] ==  pit.Displayname {   
             
            //log.Println("fffffffiiiiii", uff.Displayname)
                      
                    //serveWs(h , "@me", w , r )

            

        } else {
            http.Redirect(w, r, "/select_team/", http.StatusFound) 

        } 

      } else {
          ///http.Redirect(w, r, "/join_team", http.StatusFound) 

      } 
    
    } else { 
          w.Header()["Location"] = []string{"/login"}
	  w.WriteHeader(http.StatusTemporaryRedirect)
    }









    } else {
        log.Println("pammmmmmmmm")  
    }
    data := map[string]interface{}{
		"Host": r.Host,
    } 
    
    
      
    data["UserData"] = session.Values["username"]
    // Print secret message
    tmpl := template.Must(template.ParseFiles("assets/chat.html")) 
    tmpl.Execute(w, data) 











        }

}

func (h *Hub) SecretBis() func(http.ResponseWriter, *http.Request) {
         return func(w http.ResponseWriter, r *http.Request) {
                 segs := strings.Split(r.URL.Path, "/")
    log.Println(segs[2], "Themmm")  
    log.Println(segs[3], "Them ")  
session, _ := stori.Get(r, "cookie-name")
    if cookie, err := r.Cookie("cookie-name"); err == http.ErrNoCookie {
                log.Println(cookie)
                w.Header()["Location"] = []string{"/login"}
		w.WriteHeader(http.StatusTemporaryRedirect)

    } else if auth, ok := session.Values["species"].(bool); !ok || !auth {
    // Check if user is authenticated
    
        log.Println(session.Values["species"].(bool))  
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    } else if session.Values["username"] != nil {
    log.Println(session.Values["species"].(bool))
    user := session.Values["username"] 
    
        log.Println(user)
    } else {
        log.Println("pammmmmmmmm")  
    }
    data := map[string]interface{}{
		"Host": r.Host,
    } 
    

    

      
    data["UserData"] = session.Values["username"]
    // Print secret message
    tmpl := template.Must(template.ParseFiles("assets/chat.html")) 
    tmpl.Execute(w, data)










         } 



}




func login(w http.ResponseWriter, r *http.Request) {
    session, _ := stori.Get(r, "cookie-name")
  log.Println(session)  
//log.Println(r.Method)
    // Authentication goes here
    // ...
    err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		//w.WriteHeader(http.StatusInternalServerError)
		return
	}
    email := r.Form.Get("species")
    pass := r.Form.Get("password")
    
    if r.Method == "POST" {
//log.Println(email)
    // Set user as authenticated
    
        
       if email != "" && pass != ""{
        
    

        session.Values["species"] = true
        session.Values["username"] = email
        session.Values["password"] = pass  
        session.Save(r, w)

        usyu, erra := store.GetUserByEmail(email)
        if usyu == nil {
             log.Println("erreur", erra)
             //return
             pyt, errty := store.GetTeamMemberByEmail(email)
             log.Println("err", errty)
             //pit, errt := store.GetTeamById(pyt.Teamid)
             if pyt == nil {
                   http.Redirect(w, r, "/join_team", http.StatusFound)

             }   


      } else {
        log.Println("mmmmmmmmmmm", usyu.Email)
        } 
        http.Redirect(w, r, "/select_team/", http.StatusFound)
        

       }
     

    } else {
       


    }
   
    
  
    http.ServeFile(w, r, "assets/index.html")    
    
    
}

func (h *Hub) checkRoom(name string) *Room {
        n, present := h.roomis[name]
        if present == false {
             return nil
        } 
        return n.room

}

func (h *Hub) chackManager(name string) *managedRoom {
        n, present := h.roomis[name]
        if present == false {
             return nil
        } 
        return n

}

func (h *Hub) chack() map[string]*managedRoom {
         
        return h.roomis

}

func (h *Hub) NotFoundHandler() func(http.ResponseWriter, *http.Request) {
   return func(w http.ResponseWriter, r *http.Request) {
       w.WriteHeader(http.StatusNotFound)
       log.Println("custom 404")
   }
 
}


func (h *Hub) Handleraa() func(http.ResponseWriter, *http.Request){
       
	return func(w http.ResponseWriter, r *http.Request) {
              http.ServeFile(w, r, "assets/pam/favicon.ico")
     }

}




func (h *Hub) LoginQuat() func(http.ResponseWriter, *http.Request){
       
	return func(w http.ResponseWriter, r *http.Request) {
            vars := mux.Vars(r)  
            //_, _, lir, _ := store.GetTeamsAlli()
             fmt.Println("i am THE room", vars["room"])
              ua := r.Header.Get("User-Agent")
             fmt.Println("i am THE agent", ua)           

paz, _ := store.GetRoomsById()
log.Println("custom chat", paz)

_, _, top, zas := store.GetAlliByArray(paz)
log.Println("custom chat", zas)
poo, oor := store.GetTeams()
pao, oar := store.GetRooms()
log.Println("rooom", oor)
log.Println("rooom", oar)        
log.Println("rooom", poo)
log.Println("rooom", pao)

//h.GetTeamis(pao, poo)
//h.Teamis = top




h.Teamis = top
log.Println("custom chat", h.Teamis)
           for _, v := range h.Teamis {
              vars := mux.Vars(r) 
              fmt.Println("Hello", v)  
              for j, k := range v { 
              fmt.Println("Hello", k.teams) 
                       for _, kui := range k.teams {
                            fmt.Println("Hello",kui[0].Name)
                            
                            for _,pio := range kui {
                                fmt.Println("teamtyuuuuuuu", pio)
                                team := pio
                                if vars["team"] == team.Displayname {
                                         appTeam = team  
                                         fmt.Println("im a the team", vars["team"])  
                                         ///http.ServeFile(w, r, "assets/chat.html")
                                  } else {
                                       //w.WriteHeader(http.StatusNotFound)
                                        log.Println("custom 404")

                                  } 

                            }




                       }  
              fmt.Println("Hello", k.rooms) 
              //fmt.Println("Hellonnnnnnnnnn", k.rooms[i.Name]) 
              //fmt.Println("Hello", k.rooms[i.Name][0]) 
              //fmt.Println("Hello", k.rooms[i.Name][1])  
              fmt.Println("Hello", j)   
                       for o, ku := range k.rooms {
                            vars := mux.Vars(r) 
                            fmt.Println("team2222",o)  /// team
                            for _, kug := range ku {
                                fmt.Println("Hellonnnnnnbbbbbbbb",kug)
                                room := kug
                                
                                //fmt.Println("i am THE room", vars["room"])
                              if appTeam.Displayname == vars["team"] && vars["room"] == room.Displayname && appTeam.Id == room.Teamid {
                                    fmt.Println(vars["room"])
                                    fmt.Println("teamtyuuuuuuuggggggggggggggggggggggggggg", appTeam) 
                                    http.ServeFile(w, r, "assets/chat.html")
                              } else {
                                       ///w.WriteHeader(http.StatusNotFound)
                                        log.Println("custom 404nnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnn")

                              } 

                            }
                       }          

              }

        }

    

            
     }

}

func (h *Hub) HandlerQuat() func(http.ResponseWriter, *http.Request){
          return func(w http.ResponseWriter, r *http.Request) {

//serveWs(h , "@me", w , r )


   
     }

}

func (h *Hub) LoginTer(c Connection, author string, addr string) {
            log.Println("custom 404", c)
            log.Println("custom 404", addr)
}






func (h *Hub) Login() func(http.ResponseWriter, *http.Request){
       return func(w http.ResponseWriter, r *http.Request) {
                 params := mux.Vars(r)
	teamname := params["team_name"]
log.Println("varice", teamname)
      session, _ := stori.Get(r, "cookie-name")  
datassi := map[string]interface{}{
		"Host": r.Host,
    }             
        data := map[string]interface{}{}
datassi["UserData"] = session.Values["username"]
numi := datassi["UserData"].(string)
uff, _ := store.GetTeam("pppppppp")
log.Println("fffffffiiiiiioooooooooo", uff.Displayname)
tams, err := store.GetTeams()
        log.Println("store", tams) 
if err != nil {
log.Println("err", err) 
}
pyt, errty := store.GetTeamMemberByEmail(numi)
if errty != nil {
log.Println("err", errty) 
}
pit, errt := store.GetTeamById(pyt.Teamid)
//pit, errt := store.GetTeamByEmail(numi)
if errt != nil {
    log.Println("oooo", errt)
}
if pit != nil {
    //log.Println("handleroooo", pit.Displayname) 
    	data["UserData"] = session.Values["username"]
        
        
        if data["UserData"] == nil {
              http.Redirect(w, r, "/login", http.StatusFound)

        } else {

       datao := map[string]interface{}{
		"Host": r.Host,
    }
       datao["Teams"] = pit.Displayname

       tmpl := template.Must(template.ParseFiles("assets/indexo.html")) 
       tmpl.Execute(w, datao) 


        //http.ServeFile(w, r, "assets/indexo.html")
     
        }






} else {
http.Redirect(w, r, "/join_team", http.StatusFound)
         
}
//log.Println(TeamListToJson(tams))
       //for i := 0; i < len(tams); i++ {
		//log.Println(tams[i].Displayname)
	//}


                  

    }

}

func (h *Hub) Login2() func(http.ResponseWriter, *http.Request){
       return func(w http.ResponseWriter, r *http.Request) {
                 //segs := strings.Split(r.URL.Path, "/")
    //log.Println(segs[2], "segment2")

      session, _ := stori.Get(r, "cookie-name")               
        data := map[string]interface{}{}
        data["Password"] = session.Values["password"]
	data["UserData"] = session.Values["username"]
        numi := data["UserData"].(string)
        log.Println("handler", numi)
        

        if r.Method == "POST" {
          err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		//w.WriteHeader(http.StatusInternalServerError)
		return
	}
         team := r.Form.Get("serveur")
         
         //log.Println("lll", team) 
          ou6 := &UniqueRand{} 
        Connesiaa := ou6.generated
        if Connesiaa == nil {
          Connesiaa = make(map[int]bool)
	        ou6.generated = Connesiaa
        }
        ou6.generated[45] = true 
        pass := data["Password"].(string)
        numi := data["UserData"].(string)
        //log.Println("handler", numi)
          
        log.Println("handler", pass)    
        id := int64(ou6.Int())
	hash := CreateHash(strconv.FormatInt(id, 10), numi)
	sessionaa := &Session{
		Authori: numi,
		ID:     id,
		Hash:   hash,
	}
        now := time.Now()
        usyu, erra := store.GetUserByEmail(numi)
        log.Println("lllaaauuuuuuuu", usyu.Email)
        if erra != nil {
        log.Println("lllaaa", erra)
        } 
        appSession =  sessionaa 


        log.Println("llloooooooooooooooooooooooooo", appSession)
        log.Println("team", team)

firstTeam := &Team{Id: strconv.FormatInt(id, 10), Name: team , Description: team , Displayname: team, Email: appSession.Authori, CreateAt: now.Unix() }


        if usyu.Email != "" { 
        log.Println("taf", store.CreateTeamWithUser(firstTeam, usyu.Email))
        } else {
        password, err := GeneratePasswordHash([]byte(pass))
        if err != nil {
        log.Println("team", err) 
        }
        user := &User{Email: numi, CreatedAt: now, UpdatedAt: now, Password: pass, HashedPassword: password }
        store.CreateUser(user) 
        usyua, errau := store.GetUserByEmail(user.Email)
        log.Println("lll", usyua.Email) 
        if errau != nil {
        log.Println("lllaaa", errau)
        }

        }
        h.Teams = append(h.Teams, firstTeam)   
        //log.Println("lll", team) 
         
        } else {

        }

      http.ServeFile(w, r, "assets/indexou.html")
     
        
                  

    }

}

func (h *Hub) Loginaaaa() func(http.ResponseWriter, *http.Request){
         return func(w http.ResponseWriter, r *http.Request) {
          session, _ := stori.Get(r, "cookie-name")               
        data := map[string]interface{}{}

	data["UserData"] = session.Values["username"]
        numi := data["UserData"].(string)
         log.Println("teamuuuuuuuuuus", numi)
         tams, err := store.GetTeams()
        //log.Println("store", tams) 
if err != nil {
log.Println("err", err) 
}
datas := tams
datao := map[string]interface{}{
		"Host": r.Host,
    }
datao["Teams"] = datas

if r.Method == "POST" {
          err := r.ParseForm()
         var selected string


	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		//w.WriteHeader(http.StatusInternalServerError)
		return
	}
         if r.FormValue("teams") != "" {
		selected = r.FormValue("teams")
	}

	
         //log.Println("err", selected)  
        for _, nui := range tams {
            
           if selected ==  nui.Displayname {
//log.Println("errff", nui.Id)
memberi := &TeamMember{Teamid: nui.Id, Userid: numi, Roles: "invite"}
               store.CreateTeamMember(memberi)

            }   

        }

       put, erq := store.GetTeamMemberByEmail(numi)
log.Println("errff", erq)
log.Println("errff", put)
http.Redirect(w, r, "/serveur/channels/@me/", http.StatusFound)
       if put != nil {
          
       }  

}




tmpl := template.Must(template.ParseFiles("assets/indexa.html")) 
tmpl.Execute(w, datao)


//http.ServeFile(w, r, "assets/indexa.html")




         }

}

func (h *Hub) Handler() func(http.ResponseWriter, *http.Request){
       
	return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)        
        session, _ := stori.Get(r, "cookie-name")               
        data := map[string]interface{}{}

	data["UserData"] = session.Values["username"]
        numi := data["UserData"].(string)
        ///log.Println("teamuuuuuuuuuus", h.Teams) 

        uff, _ := store.GetTeam("pppppppp")
	log.Println("lassioo", uff.Displayname)
        //tams, err := store.GetTeams()
        //log.Println(err)
        put, erq := store.GetTeamMembersByEmail(numi)
        
        pyt, errty := store.GetTeamMemberByEmail(numi)
        if errty != nil {
        log.Println("oooo", errty)
        } 
        pit, errt := store.GetTeamById(pyt.Teamid)
        if errt != nil {  
        log.Println("oooo", errt)
        } 
        h.Members = put
        if erq != nil {
        log.Println("oooo", erq)
        }
        log.Println("handleroooomypitttttttttttttttttttttooooo", pit)
        if pit == nil {
        //log.Println(store.CreateTeamMember(pyt))
        piti, errti := store.GetTeamByEmail(numi) 
        log.Println("oooo", piti)
        if errti != nil {
        log.Println("oooo", errti)
        }

        }

        

        //memberi := &TeamMember{Teamid: pit.Id, Userid: numi, Roles: "admin"}
                    //log.Println("mem", memberi)
        
                    //log.Println(store.CreateTeamMember(memberi))

        if vars["team"] ==  pit.Displayname {   
                    
                      
                    //serveWs(h , "@me", w , r )

             

        } else if vars["team"] !=  uff.Displayname {
                //http.Redirect(w, r, "/select_team/", http.StatusFound)
                log.Println("nounou", vars)
                //log.Println("eror", err)
                log.Println("eror", vars["team"])
        } else {


        } 

        
        
        //h.HistorySess = append(h.HistorySess, sessionaa)
        
        //hi, ici := vars["name"]
        

        

    }

}

func logout(w http.ResponseWriter, r *http.Request) {
    session, err := stori.Get(r, "cookie-name")
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
    }	
    // Revoke users authentication
    session.Values["species"] = false
    session.Values["username"] = ""
    session.Options.MaxAge = -1
    err = session.Save(r, w)
    log.Println(session.Values["username"])
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
    } else {
       http.Redirect(w, r, "/login", http.StatusFound)
    }
    

}

func (h *Hub) Test() func(http.ResponseWriter, *http.Request){
           return func(w http.ResponseWriter, r *http.Request) {

segs := strings.Split(r.URL.Path, "/")
    log.Println(segs[2], "segment2")
vars := mux.Vars(r)  
log.Println(vars["team_id"], "vars")

if len(h.Sessions) == 0 {

w.WriteHeader(http.StatusNotFound)
log.Println("custom 404")

} else {

 
for _, msg := range h.Sessions {
        //log.Println("evi", msg.Authori)
	appSession = msg 
}	
        
        //data := map[string]interface{}{}
d1 := []byte(`[]`)


	//data["UserData"] = session.Values["username"]
        numi := appSession.Authori
pyt, errty := store.GetTeamMemberByEmail(numi)
log.Println("err", errty)
pit, errt := store.GetTeamById(pyt.Teamid)
log.Println("nnnnnn", errt)
log.Println("nnnnnn", pit.Id)
if vars["team_id"] == pit.Id {

 w.Header().Set("Content-Type", "application/json")
  w.Write([]byte(pit.ToJson()))
} else {
w.Header().Set("Content-Type", "application/json")
w.Write(d1)
}

}

   }
}


func (h *Hub) Test2() func(http.ResponseWriter, *http.Request){
           return func(w http.ResponseWriter, r *http.Request) {
segs := strings.Split(r.URL.Path, "/")
    log.Println(segs[2], "segment2")
vars := mux.Vars(r)  
d1 := []byte("[]")
log.Println(vars["channel_id"], "vars")
poto, oto := store.PostsForChannel(vars["channel_id"])
log.Println("nnnnnn", oto)
log.Println("nnnnnn", poto)
if len(poto) > 0 {
w.Header().Set("Content-Type", "application/json")
  w.Write([]byte(ArrayToJson(poto)))
} else {
w.Header().Set("Content-Type", "application/json")
w.Write(d1)

}



   }
}









func main() {
    //port := flag.String("p", "8000", "port to serve on")
    port := os.Getenv("PORT")
    //flag.Parse()
    
    ro := mux.NewRouter()
      
    hub := newHub()
    ou3 := &UniqueRand{} 
        Connesi := ou3.generated
        if Connesi == nil {
          Connesi = make(map[int]bool)
	        ou3.generated = Connesi
        }
        ou3.generated[12] = true 
    //go hub.run()

    hub.AddFirstRoomis()
     
    //hub.AddTeams()    
    ro.HandleFunc("/join_team" , hub.Loginaaaa())
    ro.HandleFunc("/{team_name}/display_name" , hub.Login2()).Name("display")
    //ro.HandleFunc("/channels/{room}/{name:[a-z]+}/" , hub.SecretBis()) 
     
    ro.HandleFunc("/{team_name}/" , hub.Login()).Name("team") 
     
    ro.HandleFunc("/", hub.Handleraa())
    ro.HandleFunc("/login", login)
    ro.HandleFunc("/logout", logout)  
    
     
    



    ro.HandleFunc("/{team}/channels/{room}/" , hub.LoginQuat())  
            //ro.HandleFunc("/{team}/channels/{room}", hub.HandlerQuat())   
    //ro.HandleFunc("/channels/{room}/{name:[a-z]+}", hub.Handler())

    //ro.HandleFunc("/{team}/channels/{room}", hub.Handler())
                    
     
        //ro.HandleFunc("/user", createUserHandler).Methods("POST")
	// Declare the static file directory and point it to the
	// directory we just made
	staticFileDirectory := http.Dir("./assets/")
	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for
	// "./assets/assets/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	ro.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

    					 
     			  
    //cuuu := make(chan []Room) 
    //connString := "user=tuodvhxnybdiys dbname=d7ek26bkp6ob11 password=a35c916f2d0e58c80115a2254961bd038fa4b3b79312b821091b7160f32b189d"
	//db, err := sql.Open("tuodvhxnybdiys", connString)
          
    db, err := sql.Open("postgres", "postgres://tuodvhxnybdiys:a35c916f2d0e58c80115a2254961bd038fa4b3b79312b821091b7160f32b189d@ec2-34-250-16-127.eu-west-1.compute.amazonaws.com:5432/d7ek26bkp6ob11")
	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {

		panic(err)
	}
        
        //now := time.Now()
        //password, err := GeneratePasswordHash([]byte("froggy25")); 
        //teston := append(h.HistorySess, session)
	InitStore(&dbStore{db: db})
        
        //_, _, lir, _ := store.GetTeamsAlli()
paz, _ := store.GetRoomsById()
log.Println("custom chat", paz)

poo, oor := store.GetTeams()
pao, oar := store.GetRooms()
log.Println("rooom", oor)
log.Println("rooom", oar)
log.Println("rooom", poo)
log.Println("rooom", pao)


poum, pouma, top, zas := store.GetAlliByArray(paz)
log.Println("custom chat", zas)
log.Println("custom chat", pouma)
log.Println("custom chat", poum)
hub.Teamis = top

input := []byte("I just love bold text.")
output := blackfriday.MarkdownCommon(input)

fmt.Println(string(output))
fmt.Println(string(input))

emoji :=  &Emoji{Id: autoId(), Keyi: "1F493", Valeur: "😿", Descriptor: "Beating Heart"  }
rmajis, errtoi := store.GetEmojisAll()
log.Println("erreur", errtoi)
log.Println("erreur", rmajis)

var x = '😂'
fmt.Printf("%d\n", x)
fmt.Printf("%+q", x)
log.Println("fffffff", "\u0032\uFE0F\u20E3")
//fmt.Println([]byte(emoji.ToJson()))
resy := Emoji{}

fmt.Println(string([]byte(emoji.Valeur)))
test := []byte(emoji.ToJson())
log.Println("rooom", string(test))
//byt := []byte(string(test))
//var dat map[string]interface{}
erruu := json.Unmarshal([]byte(string(test)), &resy)
log.Println("rooomiiiiiiiiiiii", erruu)
log.Println("rooomiiiiiiiiiiii", &resy)


        paii, poiu, ertyui := store.GetTeamsAlla()
log.Println("rooom", ertyui)
log.Println("rooom", paii)
log.Println("rooom", poiu)

count, zert := store.GetCountRooms()
log.Println("rooom", zert)
log.Println("coun", count)
log.Println("rasli", hub.GetUrls(paii, poiu))
log.Println("rasle", hub.GetTeamis(paii, poiu))
//log.Println("ggggggiiiiiiiiiii", hub.GetTeamis(pao, poo))
//log.Println("gggggg", hub.GetTastis(hub.GetTeamis(pao, poo)).team)
//hub.GetTeamis(pao, poo)

     //InitStores(storaa)
        ro.HandleFunc("/api/v4/teams/{team_id:[A-Za-z0-9]+}/channels", hub.Test()) 
        //listo := &ListTeam{}
        ro.HandleFunc("/api/v4/channels/{channel_id:[A-Za-z0-9]+}/posts", hub.Test2()) 
        //store.init() 
//user := &User{Email: "arduino2501@pom.fr", CreatedAt: now, UpdatedAt: now, Password: "froggy25", HashedPassword: password }
        ro.HandleFunc("/{team}/channels/{room}", func(w http.ResponseWriter, r *http.Request) {
                



		if ws, err := NewWebSocket(hub, w, r); err == nil {
                        vars := mux.Vars(r)  
                        log.Println(vars["team"])
                        

                        
                        for _, v := range hub.Teamis {
                        vars := mux.Vars(r) 
                        fmt.Println("Hellotouuttttttttttttt", v)  
              for j, k := range v { 
              fmt.Println("Hello k teams ", k.teams) 
                       for _, kui := range k.teams {
                            fmt.Println("server",kui[0].Name)
                            
                            
                            for _,pio := range kui {
                                fmt.Println("teamtyuuuuuuu", pio)
                                team := pio
                                if vars["team"] == team.Displayname {
                                         fmt.Println("mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm")
                                  } 

                            }
                            
                            


                       }  
              fmt.Println("Hello k rooms", k.rooms) 
              //fmt.Println("Hellonnnnnnnnnn", k.rooms[i.Name]) 
              //fmt.Println("Hello", k.rooms[i.Name][0]) 
              //fmt.Println("Hello", k.rooms[i.Name][1])  
              fmt.Println("Hello 2eme boucle", j)   
                       for o, ku := range k.rooms {
                            vars := mux.Vars(r) 
                            fmt.Println("team2222",o)  /// team
                            for _, kug := range ku {
                                fmt.Println("i m a room",kug)
                                room := kug
                                
                                if vars["room"] == room.Displayname && appTeam.Displayname == vars["team"]{
                                      fmt.Println("i am THEEEEEEEEEEEEEEEEEEEEEEEEE room", vars["room"])
                                      log.Println("appTeam", appTeam)
                                      hub.Register(room, ws)  
                                 

                                
                                 
                         			addr := strings.Split(r.RemoteAddr, ":")[0]

			loginUser := func(uname string) {
                                room.Author = &Author{
					ID:       ou3.Int(),
					Color:    UtilGetRandomColor(),
					Username: uname,
                                        Room: room,  
				}
                                hub.Sackets[ws] = &Author{
					ID:       ou3.Int(),
					Color:    UtilGetRandomColor(),
					Username: uname,
				}


                               


				hub.Sockets[ws] = room
                                //hub.Posts = 
                                poto, oto := store.PostsForChannel(room.Id)
                                log.Println("rooomiiiiiiiiiiiiOOOOOOOO", oto)
                                //log.Println("rooomiiiiiiiiiiiiOOOOOOOOOOO", poto)
                                hub.Posts = poto 
                              
                                
				hub.Loginnn(ws, hub.Sockets[ws].Room.Author, addr)
				hub.Users[hub.Sockets[ws].Room.Author.Username] = hub.Sockets[ws].Room.Author
	                        //hub.AppendHistoryPosts(room)
                                log.Println("usersiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii", hub.Teamis["mykeys"]["key"])
                                log.Println("users", hub.Users)
                                
				go func() {
					ws.send <- (&Event{
						Name: "clientConnect",
						Data: map[string]interface{}{
							"author":   hub.Sockets[ws],
							"nclients": len(hub.Sockets),
							"clients":  hub.Users,
                                                        "test": hub.GetTeamis(paii, poiu),
                                                        "tosts": paii,
                                                        "tasto": hub.GetUrls(paii, poiu), 
                                                        "pals": vars["team"], 
                                                        "palis": room, 
                                                        "emo": rmajis, 
                                                        "roomteams": poiu,  
                                                        "hposts":  hub.Posts,   
							"history":  hub.History,
						},
					}).Raw()
				}()
				hub.Broadcast((&Event{
					Name: "connected",
					Data: map[string]interface{}{
						"author":   hub.Sockets[ws],
                                                "room": room,  
                                                "nclients": len(hub.Sockets),
						"clients":  hub.Users,
					},
				}).Raw(), (room ) )
			}

			

			// USERNAME INPUT EVENT
			// -> Checks if name is not connected yet
			//    -> else send 'connect_reject' event
			// -> Broadcast to all clients that user has connected
			ws.SetHandler("login", func(event *Event) {
				dataMap := event.Data.(map[string]interface{})
				uname := dataMap["username"].(string)
				passwd := dataMap["password"].(string)
                                log.Println("pass", passwd)
				
				loginUser(uname)
			})

                        laginUser := func(uname string) {
                                log.Println("pass", uname)
                        }

                        ws.SetHandler("suscribe", func(event *Event) {
				dataMap := event.Data.(map[string]interface{})
				uname := dataMap["username"].(string)
				//log.Println("ccc", req.name) 
                                log.Println("pass", uname)
				log.Println("passoooooooooo", hub.Sackets[ws].Username)
                                pi, _ := store.GetTeamMemberByEmail(hub.Sackets[ws].Username)
                                log.Println("passoooooooooo", pi.Roles)
                                log.Println("passoooooooooo", pi.Userid)
                                s := strconv.FormatInt(count + 1, 10)
                                room := &Room{Id: s , Name: uname, Teamid: pi.Teamid , Displayname: uname , Authorid: pi.Userid } 
                                log.Println("passoooooooooo", pi.Teamid)
                                rat := store.CreateRoomWithUser(room , pi.Userid, pi.Teamid) 
                                log.Println("passoooooooooo", rat)


				laginUser(uname)
			})

			// CHAT MESSAGE EVENT
			// -> Attach username to message
			// -> Broadcast the chat message to all users
			ws.SetHandler("message", func(event *Event) {
				if len(strings.Trim(event.Data.(string), " \t")) < 1 {
					return
				}
				author := hub.Sackets[ws]
                                m := int64(author.ID)
                                log.Println("ibbiiiiiii", m)
                                //if hub.TempHistoryLength(m) > 10 {
					//go func() {
						//ws.Out <- (&Event{
							//Name: "spamTimeout",
							//Data: nil,
						//}).Raw()
					//}()
					//return
			        //}
                         

                                mass := &Message{ Content:   strings.Replace(event.Data.(string), "\\n", "\n", -1),  }
    bi := []byte(event.Data.(string))  
     yi := make([]byte, 0)
       for _, b := range bi {
           if escaped(b) > -1 {
              
             //return bi
           } else {
           
            log.Println("ibbiiiiiii", b)
            yi = append(yi, b)
            mass.Content = string(yi)
           }
     }
    

log.Println("strip", strip_tags(mass.Content))
log.Println("authoriiiiiiiiiiiiiiiiiiiiiiiiiiiiii", event.Data.(string))
log.Println("authoriiiiiiiiiiiiiiiiiiiiiiiiiiiiii", testEsc(event.Data.(string)))
			        
				event.Data = &Message{
					Author:    author.Username,
                                        Color:     author.Color,
                                        Room:      room.Displayname,
					Content:   string(blackfriday.MarkdownCommon([]byte(mass.Content))),
					Timestamp: time.Now().Unix(),
					Id:        autoId(),
				}
                                //log.Println("ggggg", room)
				hub.Broadcast(event.Raw(), ( room ))
post := &Post{Id: autoId(), Createat: time.Now().Unix(), Userid: author.Username, Channelid: room.Id, Rootid: strconv.FormatInt(m, 10),  Parentid: strconv.FormatInt(m, 10), Message: strip.StripTags(mass.Content), Color: author.Color, Typo: "message" , Props: &Props{
        Username: author.Username,
        Markdown: mass.Content,
        Channels: struct {
            Channelid string `json:"channelid,omitempty"`
            
        }{
            Channelid:   room.Id,
            
        },
    } , }   
    
        heelo := store.CreatePost(post)
log.Println("post", post)
log.Println("posti", heelo)

                                 
				//hub.AppendHistory(event)
                                //mi := int64(author.ID) 
				//hub.EnqueueTempHistory(mi)
			})

			ws.SetHandler("deleteMessage", func(event *Event) {
                                //author := hub.Sackets[ws]
                                //m := int64(author.ID)
				data := event.Data.(map[string]interface{})
                                n, err := strconv.ParseInt(data["msgid"].(string), 10, 64)
                                if err == nil {
                                //fmt.Printf("%d of type %T", n, n)
                                }
				msgID := n
                                //log.Println("deletennnnnnnnn", hub.DeletePostByID(msgID))
				if msg := hub.DeletePostByID(msgID); msg != nil {
					        heelo := store.DeletePostsForChannel(data["msgid"].(string), room.Id)
                                                log.Println("deletennnnnnnnn", heelo)
						eventOut := &Event{
							Name: "messageDeleted",
							Data: msg,
						}
						hub.Broadcast(eventOut.Raw(), (room ))
					
				}
			})

                        ws.SetHandler("editMessage", func(event *Event) {
                            data := event.Data.(map[string]interface{})
                                log.Println("editnnnnnnnnnnnnn", data["mycontent"].(string))
                                log.Println("editnnnnnnnnnnnnn", data["msgid"].(string))
                                heelo := store.UpdatePostsForChannel(data["msgid"].(string), room.Id, data["mycontent"].(string) )
                                log.Println("update", heelo)


                        })  
			// DISCONNECT EVENT
			// -> Broadcast to all clients that
			//    user has disconnected
			ws.SetHandler("disconnected", func(event *Event) {
				dataMap := event.Data.(map[string]interface{})
				uname := dataMap["name"].(string)
				delete(hub.Users, uname)
				dataMap["clients"] = hub.Users
				event.Data = dataMap
				hub.Broadcast(event.Raw(), (room ))
			})








                              if vars["room"] == room.Displayname {

                              }

                            }
                        }   

                    }          

              } 















                        }
			
		}


	})
        

      
    
    
    //http.Handle("/channel/" +name.Name , hub)
       
    
    //log.Println("serving...:", *port)  
    log.Println("serving...:", port)  
    http.ListenAndServe(":"+port, ro)
    
    
}
