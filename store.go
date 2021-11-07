package main

// The sql go library is needed to interact with the database
import (
          "log"
          "context"  
          "database/sql"
          "github.com/lib/pq"
          "github.com/jinzhu/gorm"
	  "github.com/pkg/errors"
          "database/sql/driver"
          "encoding/json"
)

var (
	ctx context.Context
	db  *sql.DB
        dbtost  []*Team
    
)


type StoreResult struct {
	Data interface{}
	
}

type StoreChannel chan StoreResult

func Must(sc StoreChannel) interface{} {
	r := <-sc
	

	return r.Data
}
// Our store will have two methods, to add a new bird,
// and to get all existing birds
// Each method returns an error, in case something goes wrong
type Stora interface {
     Team() TeamStore   

}

type Store interface {
	//CreateBird(bird *Bird) error
        CreateUser(user *User) error
        CreateEmoji(e *Emoji) error
        CreatePost(post *Post) error
        CreateTableTeams() error
        CreateTeam(ti *Team) error
        CreateSession(sess *Session) error 
        CreateRoomWithUser(room *Room, email string, tom string) error
        GetCountRooms() (int64 , error)
        GetRoomsById() ([]string, error)
        GetTeamsAll() ([]*Team, error)
        GetRoomsAll() (map[string]*Room, error)
        GetEmojisAll() ([]*Emoji, error)

        DeletePostsForChannel(id string, cid string) error 
        UpdatePostsForChannel(id string, cid string, content string) error
        //GetTeamsAlli() ([]*Room, []*Team, map[string]*managedTeam, error)
        GetTeamsAlla() ([]*Room, []*Team, error)
        GetSession(hash string) (*Session, error)
        GetSessionByName(name string) (*Session, error) 
        GetAlliByArray(codes []string) ([]*Room, []*Team, map[string]map[string]*managedTeam, error)

        CreateTeamMember(member *TeamMember) error
        GetTeamMemberByEmail(email string) (*TeamMember, error) 
        GetTeamMembersByEmail(email string) ([]*TeamMember, error)
        GetTeamById(id string) (*Team, error)
        GetMembersAll(email string) ([]*Team, error)

        getParentsPosts(channelId string) StoreChannel

        PostsForChannel(cid string) ([]*Post, error)

        GetRoomsByTeamId(id string, email string) (*Room, error)
        GetRoomByEmail(email string) (*Room, error)
        GetRoomByAuthor() (*Room, error)
        CreateRoom(room *Room) error
        CreateTeamWithUser(ti *Team, email string) error
        GetPosts() ([]*Post, error)
        //GetPostsByRoom(post *Post) (*Post, error) 
	//GetBirds() ([]*Bird, error)
        GetRoomTest(codes []string) (*Room, error)
        GetTeams() ([]*Team, error)
        GetRooms() ([]*Room, error)
        GetTeamsByEmail(email string) ([]*Team, error) 
        GetUsers() ([]*User, error)
        CreateTableUser() error
        GetUserByEmail(email string) (*User, error)
        GetTeamByEmail(email string) (*Team, error)
        GetTeam(namee string) (*Team, error)
}
type TeamStore interface {
	Save(team *Team) StoreChannel
}
// The `dbStore` struct will implement the `Store` interface
// It also takes the sql DB connection object, which represents
// the database connection.
type dbStore struct {
	db *sql.DB
        dbi *gorm.DB
         
}


type Props struct {
    Username    string   `json:"username,omitempty"`
    Markdown    string   `json:"markdown,omitempty"` 
    Channels  struct {
        Channelid string `json:"channelid,omitempty"`
    } `json:"channels,omitempty"`
}

func (p Props) Value() (driver.Value, error) {
    return json.Marshal(p)
}

func (a *Props) Scan(value interface{}) error {
    b, ok := value.([]byte)
    if !ok {
        return errors.New("type assertion to []byte failed")
    }

    return json.Unmarshal(b, &a)
}



func (t *Team) Save(Team) StoreChannel { 
      resi := StoreResult{t}
      storeChannel := make(StoreChannel, 1)
	go func() {
		result := resi
		
		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel 
}


func (store *dbStore) CreateTableTeams() error {
const addTeamSQL = `
			CREATE TABLE public.teams(
 				id VARCHAR(26) PRIMARY KEY,
                                UpdateAt bigint,
                                CreateAt bigint,
                                DeleteAt bigint,
                                DisplayName VARCHAR(64),
                                Namee VARCHAR(64) UNIQUE,
                                Email VARCHAR(128),
 				Typee VARCHAR(255),
                                InviteId VARCHAR(32)
                                
			);
		`

		_, err:= store.db.Exec(addTeamSQL)
		return err

}




func (store *dbStore) CreateTableUser() error {
const addUserSQL = `
			CREATE TABLE public.users(
 				id serial PRIMARY KEY,
 				email text UNIQUE NOT NULL,
				hashed_password bytea NOT NULL,
 				created_at TIMESTAMP NOT NULL,
 				updated_at TIMESTAMP NOT NULL,
 				deleted_at TIMESTAMP
			);
		`

		_, err:= store.db.Exec(addUserSQL)
		return err



}


func (store *dbStore) CreateEmoji(e *Emoji) error {
	// 'Bird' is a simple struct which has "species" and "description" attributes
	
 _, err := store.db.Query("INSERT INTO emoji(id, keyi, valeur , descriptor) VALUES ($1,$2,$3,$4)", e.Id, e.Keyi, e.Valeur, e.Descriptor)
 return err
}



func (store *dbStore) CreateTeamMember(member *TeamMember) error {
	// 'Bird' is a simple struct which has "species" and "description" attributes
	
 _, err := store.db.Query("INSERT INTO team_members(teamid, userid , roles) VALUES ($1,$2,$3)", member.Teamid, member.Userid, member.Roles)
 return err
}

func (store *dbStore) CreateSession(sess *Session) error {
	// 'Bird' is a simple struct which has "species" and "description" attributes
	
 _, err := store.db.Query("INSERT INTO sessions(id, authori, hash) VALUES ($1,$2,$3)", sess.ID, sess.Authori, sess.Hash)
 return err
}



func (store *dbStore) CreateRoom(room *Room) error {
	// 'Bird' is a simple struct which has "species" and "description" attributes
	
 _, err := store.db.Query("INSERT INTO rooms(id, name , teamid, createat) VALUES ($1,$2,$3)", room.Id, room.Teamid, room.Name, room.CreateAt)
 return err
}



func (store *dbStore) CreateUser(user *User) error {
	// 'Bird' is a simple struct which has "species" and "description" attributes
	
 _, err := store.db.Query("INSERT INTO users(email, hashed_password , created_at, updated_at) VALUES ($1,$2,$3,$4)", user.Email, user.HashedPassword, user.CreatedAt, user.UpdatedAt)
 return err
}

func (store *dbStore) CreateTeam(ti *Team) error {
	// 'Bird' is a simple struct which has "species" and "description" attributes
	
 _, err := store.db.Query("INSERT INTO teams(id, namee, email, createat, displayname) VALUES ($1,$2,$3,$4,$5)", ti.Id, ti.Name, ti.Email, ti.CreateAt, ti.Displayname)
 return err
}

func (store *dbStore) GetTeam(namee string) (*Team, error) {
	var team Team
	

    dbi, err := gorm.Open("postgres", "postgres://tuodvhxnybdiys:a35c916f2d0e58c80115a2254961bd038fa4b3b79312b821091b7160f32b189d@ec2-34-250-16-127.eu-west-1.compute.amazonaws.com:5432/d7ek26bkp6ob11")
        //connString := "user=tuodvhxnybdiys dbname=d7ek26bkp6ob11 password=a35c916f2d0e58c80115a2254961bd038fa4b3b79312b821091b7160f32b189d"
        //dbi, err := gorm.Open("tuodvhxnybdiys", connString) 
        defer dbi.Close() 
        if err != nil {

		log.Println(err)
	}

	if err := dbi.First(&team, Team{namee: namee}).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, err
		}
		return nil, errors.Wrap(err, "unable to get team")
	}

	return &team, nil
}

func (store *dbStore) CreateRoomWithUser(room *Room, email string, tom string) error {
	// 'Bird' is a simple struct which has "species" and "description" attributes
	
 _, err := store.db.Query("INSERT INTO rooms(id, name , teamid, authorid) VALUES ($1,$2,$3,$4)", room.Id, room.Name, tom, email)
 return err
}



func (store *dbStore) CreateTeamWithUser(ti *Team, email string) error {
	// 'Bird' is a simple struct which has "species" and "description" attributes
	
 _, err := store.db.Query("INSERT INTO teams(id, namee, email, createat, displayname) VALUES ($1,$2,$3,$4,$5)", ti.Id, ti.Name, email, ti.CreateAt, ti.Displayname)
 return err
}


func (store *dbStore) CreatePost(post *Post) error {
//idoString := StringInterface{ "username": "mon nom", }["username"].(string)
	// 'Bird' is a simple struct which has "species" and "description" attributes


 _, err := store.db.Query("INSERT INTO posts(id, createat, userid, channelid, rootid, parentid, message, color, typo, props) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)", post.Id, post.Createat, post.Userid, post.Channelid, post.Rootid, post.Parentid, post.Message, post.Color, post.Typo, post.Props)
 return err
}


func (store *dbStore) GetTeamById(id string) (*Team, error) {
	var team Team
        //connString := "user=tuodvhxnybdiys dbname=d7ek26bkp6ob11 password=a35c916f2d0e58c80115a2254961bd038fa4b3b79312b821091b7160f32b189d"
        //dbi, err := gorm.Open("tuodvhxnybdiys", connString) 
       

    dbi, err := gorm.Open("postgres", "postgres://tuodvhxnybdiys:a35c916f2d0e58c80115a2254961bd038fa4b3b79312b821091b7160f32b189d@ec2-34-250-16-127.eu-west-1.compute.amazonaws.com:5432/d7ek26bkp6ob11")
        defer dbi.Close()
        if err != nil {

		log.Println(err)
	}

	if err := dbi.First(&team, Team{Id: id}).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "unable to get teamuuu")
	}

	return &team, nil
}



func (store *dbStore) GetTeamByEmail(email string) (*Team, error) {
	var team Team
        //connString := "user=tuodvhxnybdiys dbname=d7ek26bkp6ob11 password=a35c916f2d0e58c80115a2254961bd038fa4b3b79312b821091b7160f32b189d"
        //dbi, err := gorm.Open("tuodvhxnybdiys", connString) 
        dbi, err := gorm.Open("postgres", "postgres://tuodvhxnybdiys:a35c916f2d0e58c80115a2254961bd038fa4b3b79312b821091b7160f32b189d@ec2-34-250-16-127.eu-west-1.compute.amazonaws.com:5432/d7ek26bkp6ob11")
        defer dbi.Close()
        if err != nil {

		log.Println(err)
	}

	if err := dbi.First(&team, Team{Email: email}).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "unable to get teamuuu")
	}

	return &team, nil
}

func (store *dbStore) GetTeamMemberByEmail(email string) (*TeamMember, error) {
	var member TeamMember
	
    dbi, err := gorm.Open("postgres", "postgres://tuodvhxnybdiys:a35c916f2d0e58c80115a2254961bd038fa4b3b79312b821091b7160f32b189d@ec2-34-250-16-127.eu-west-1.compute.amazonaws.com:5432/d7ek26bkp6ob11")
        //connString := "user=tuodvhxnybdiys dbname=d7ek26bkp6ob11 password=a35c916f2d0e58c80115a2254961bd038fa4b3b79312b821091b7160f32b189d"
        //dbi, err := gorm.Open("tuodvhxnybdiys", connString) 
        defer dbi.Close()
        if err != nil {

		log.Println(err)
	}

	if err := dbi.First(&member, TeamMember{Userid: email}).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "unable to get teamuuuiiiiiiiii")
	}

	return &member, nil
}



func (store *dbStore) GetUserByEmail(email string) (*User, error) {
	var user User
        //connString := "user=tuodvhxnybdiys dbname=d7ek26bkp6ob11 password=a35c916f2d0e58c80115a2254961bd038fa4b3b79312b821091b7160f32b189d"
        //dbi, err := gorm.Open("tuodvhxnybdiys", connString) 
        
    dbi, err := gorm.Open("postgres", "postgres://tuodvhxnybdiys:a35c916f2d0e58c80115a2254961bd038fa4b3b79312b821091b7160f32b189d@ec2-34-250-16-127.eu-west-1.compute.amazonaws.com:5432/d7ek26bkp6ob11")
    
        defer dbi.Close()
        if err != nil {

		log.Println(err)
	}

	if err := dbi.First(&user, User{Email: email}).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "unable to get user")
	}

	return &user, nil
}


func (store *dbStore) GetRoomsAll() (map[string]*Room, error) {
rows, err := store.db.Query("SELECT id, teamid, displayname, authorid from rooms")
if err != nil {
		return nil, err
	}
defer rows.Close()

rooms := make(map[string]*Room) 
for rows.Next() {
		// For each row returned by the table, create a pointer to a bird,
		room := &Room{}
		// Populate the `Species` and `Description` attributes of the bird,
		// and return incase of an error
		if err := rows.Scan(&room.Id, &room.Teamid, &room.Displayname, &room.Authorid); err != nil {
			return nil, err
		}
		// Finally, append the result to the returned array, and repeat for
		// the next row
                rooms[room.Teamid] = room
		//rooms = append(rooms, room)

	}
	return rooms, nil

}

func (store *dbStore) getParentsPosts(channelId string) StoreChannel {
	storeChannel := make(StoreChannel, 1)

	go func() {
		result := StoreResult{}

		var posts []*Post
		_, err := store.db.Query(
			`SELECT q2.* FROM posts q2 INNER JOIN (SELECT DISTINCT q3.rootid FROM (SELECT rootid FROM posts WHERE channelid = :ChannelId ORDER BY createat DESC) q3
			    WHERE q3.rootid != '') q1
			    ON q1.rootid = q2.Id OR q1.rootid = q2.rootid
			WHERE
			    ChannelId = :ChannelId2
			        
			ORDER BY createat`,
			map[string]interface{}{"ChannelId": channelId})
		if err != nil {
			return 
		} else {
			result.Data = posts
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (store *dbStore) UpdatePostsForChannel(id string, cid string, content string) error {
 
 _, err := store.db.Query("UPDATE posts SET message = $3 WHERE id = $1 AND channelid = $2", id, cid, content)
 return err


}




func (store *dbStore) DeletePostsForChannel(id string, cid string) error {
 
 _, err := store.db.Query("DELETE FROM posts WHERE id = $1 AND channelid = $2", id, cid)
 return err


}

func (store *dbStore) PostsForChannel(cid string) ([]*Post, error){
rows, err := store.db.Query("SELECT id, createat, userid, channelid, rootid, parentid, message, color, typo, props from posts WHERE channelid = $1 ORDER BY createat ASC", cid) 
if err != nil {
		return nil, err
	}
defer rows.Close()
posts := []*Post{}
	for rows.Next() {
		// For each row returned by the table, create a pointer to a bird,
		post := &Post{}
                
                 
		// Populate the `Species` and `Description` attributes of the bird,
		// and return incase of an error
		if err := rows.Scan(&post.Id, &post.Createat, &post.Userid, &post.Channelid, &post.Rootid, &post.Parentid , &post.Message, &post.Color, &post.Typo, &post.Props); err != nil {
			return nil, err
		}
		// Finally, append the result to the returned array, and repeat for
		// the next row
		posts = append(posts, post)

	}
	return posts, nil


}


func (store *dbStore) GetTeamsAlla() ([]*Room, []*Team, error) {
query := ` SELECT  r.id, r.teamid, r.displayname, t.id, t.displayname
    FROM rooms AS r 
    JOIN teams AS t ON r.teamid = t.id`
rows, err := store.db.Query(query)
if err != nil {
		return nil, nil, err
	}
defer rows.Close()
roomteams := []*Team{}
rams := []*Room{}
for rows.Next() {
    room := &Room{}
    team := &Team{}
    err = rows.Scan(
      &room.Id,
      &room.Teamid,
      &room.Displayname,  
      &team.Id,
      &team.Displayname,
      
    )
    if err != nil {
		return nil, nil, err
	}
    roomteams = append(roomteams, team)
    rams = append(rams, room)
    
  }
        return rams, roomteams, nil
       
}

func (store *dbStore) GetEmojisAll() ([]*Emoji, error) {
rows, err := store.db.Query("SELECT id, keyi, valeur, descriptor from emoji ")
if err != nil {
		return nil, err
	}
defer rows.Close()
 
emojis := []*Emoji{}
for rows.Next() {
		// For each row returned by the table, create a pointer to a bird,
		emoji := &Emoji{}
                 
		// Populate the `Species` and `Description` attributes of the bird,
		// and return incase of an error
		if err := rows.Scan(&emoji.Id, &emoji.Keyi, &emoji.Valeur, &emoji.Descriptor); err != nil {
			return nil, err
		}
                
                
		// Finally, append the result to the returned array, and repeat for
		// the next row
		
                emojis = append(emojis, emoji)

	}
	return emojis, nil

}




func (store *dbStore) GetTeamsAll() ([]*Team, error) {
rows, err := store.db.Query("SELECT id, displayname, email from teams ")
if err != nil {
		return nil, err
	}
defer rows.Close()
 
teams := []*Team{}
for rows.Next() {
		// For each row returned by the table, create a pointer to a bird,
		team := &Team{}
                 
		// Populate the `Species` and `Description` attributes of the bird,
		// and return incase of an error
		if err := rows.Scan(&team.Id, &team.Displayname, &team.Email); err != nil {
			return nil, err
		}
                
                
		// Finally, append the result to the returned array, and repeat for
		// the next row
		
                teams = append(teams, team)

	}
	return teams, nil

}





func (store *dbStore) GetUsers() ([]*User, error) {
rows, err := store.db.Query("SELECT email from users")
if err != nil {
		return nil, err
	}
defer rows.Close()
users := []*User{}
for rows.Next() {
		// For each row returned by the table, create a pointer to a bird,
		user := &User{}
		// Populate the `Species` and `Description` attributes of the bird,
		// and return incase of an error
		if err := rows.Scan(&user.Email); err != nil {
			return nil, err
		}
		// Finally, append the result to the returned array, and repeat for
		// the next row
		users = append(users, &User{Email: user.Email})

	}
	return users, nil

}

func (store *dbStore) GetPosts() ([]*Post, error) {
	// Query the database for all birds, and return the result to the
	// `rows` object
	rows, err := store.db.Query("SELECT id, createat, userid, channelid, rootid, parentid, message, color, typo, props from posts")
	// We return incase of an error, and defer the closing of the row structure
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create the data structure that is returned from the function.
	// By default, this will be an empty array of birds
	posts := []*Post{}
	for rows.Next() {
		// For each row returned by the table, create a pointer to a bird,
		post := &Post{}
                
                 
		// Populate the `Species` and `Description` attributes of the bird,
		// and return incase of an error
		if err := rows.Scan(&post.Id, &post.Createat, &post.Userid, &post.Channelid, &post.Rootid, &post.Parentid, &post.Message, &post.Color, &post.Typo, &post.Props); err != nil {
			return nil, err
		}
		// Finally, append the result to the returned array, and repeat for
		// the next row
		posts = append(posts, post)

	}
	return posts, nil
}


func (store *dbStore) GetRooms() ([]*Room, error) {
rows, err := store.db.Query("SELECT id, teamid, displayname from rooms")
if err != nil {
		return nil, err
	}
defer rows.Close()
rooms := []*Room{}
for rows.Next() {
		// For each row returned by the table, create a pointer to a bird,
		room := &Room{}
		// Populate the `Species` and `Description` attributes of the bird,
		// and return incase of an error
		if err := rows.Scan(&room.Id ,&room.Teamid, &room.Displayname); err != nil {
			return nil, err
		}
		// Finally, append the result to the returned array, and repeat for
		// the next row
		rooms = append(rooms, room)

	}
	return rooms, nil

}

func (store *dbStore) GetRoomsById() ([]string, error) {
rows, err := store.db.Query("SELECT teamid, displayname from rooms")
if err != nil {
		return nil, err
	}
defer rows.Close()
rooms := []string{}
for rows.Next() {
		// For each row returned by the table, create a pointer to a bird,
		room := &Room{}
		// Populate the `Species` and `Description` attributes of the bird,
		// and return incase of an error
		if err := rows.Scan(&room.Teamid, &room.Displayname); err != nil {
			return nil, err
		}
		// Finally, append the result to the returned array, and repeat for
		// the next row
		rooms = append(rooms, room.Teamid)

	}
	return rooms, nil

}


func (store *dbStore) GetTeams() ([]*Team, error) {
rows, err := store.db.Query("SELECT id, displayname from teams")
if err != nil {
		return nil, err
	}
defer rows.Close()
teams := []*Team{}
for rows.Next() {
		// For each row returned by the table, create a pointer to a bird,
		team := &Team{}
		// Populate the `Species` and `Description` attributes of the bird,
		// and return incase of an error
		if err := rows.Scan(&team.Id, &team.Displayname); err != nil {
			return nil, err
		}
		// Finally, append the result to the returned array, and repeat for
		// the next row
		teams = append(teams, team)

	}
	return teams, nil

}

func (store *dbStore) GetTeamsByEmail(email string) ([]*Team, error) {
rows, err := store.db.Query("SELECT displayname from teams")
if err != nil {
		return nil, err
	}
defer rows.Close()
teams := []*Team{}
for rows.Next() {
		// For each row returned by the table, create a pointer to a bird,
		team := &Team{Email: email}
		// Populate the `Species` and `Description` attributes of the bird,
		// and return incase of an error
		if err := rows.Scan(&team.Displayname); err != nil {
			return nil, err
		}
		// Finally, append the result to the returned array, and repeat for
		// the next row
		teams = append(teams, &Team{Displayname: team.Displayname})

	}
	return teams, nil

}

func (store *dbStore) GetMembersAll(email string) ([]*Team, error) {
rows, err := store.db.Query("SELECT displayname FROM teams LEFT JOIN team_members ON teams.id=team_members.teamid WHERE team_members.teamid = $1", email)
if err != nil {
		return nil, err
	}
defer rows.Close()
teams := []*Team{}
for rows.Next() {
		// For each row returned by the table, create a pointer to a bird,
		team := &Team{}
		// Populate the `Species` and `Description` attributes of the bird,
		// and return incase of an error
                

		if err := rows.Scan(&team.Displayname); err != nil {
			return nil, err
		}
		// Finally, append the result to the returned array, and repeat for
		// the next row
		teams = append(teams, team)

	}
	return teams, nil


}


func (store *dbStore) GetTeamMembersByEmail(email string) ([]*TeamMember, error) {
rows, err := store.db.Query("SELECT teamid, userid, roles  from team_members where userid = $1", email)
if err != nil {
		return nil, err
	}
defer rows.Close()
teams := []*TeamMember{}
for rows.Next() {
		// For each row returned by the table, create a pointer to a bird,
		team := &TeamMember{}
		// Populate the `Species` and `Description` attributes of the bird,
		// and return incase of an error
                

		if err := rows.Scan(&team.Teamid, &team.Userid, &team.Roles); err != nil {
			return nil, err
		}
		// Finally, append the result to the returned array, and repeat for
		// the next row
		teams = append(teams, team)

	}
	return teams, nil

}


func (store *dbStore) GetAlliByArray(codes []string) ([]*Room, []*Team, map[string]map[string]*managedTeam, error) {
query := ` SELECT r.id, r.teamid, r.displayname, t.id, t.displayname
    FROM rooms AS r, teams AS t WHERE r.teamid = any($1) AND r.teamid = t.id`
rows, err := store.db.Query(query, pq.Array(codes))
if err != nil {
		return nil, nil, nil, err
	}
defer rows.Close()
y := make(map[string][]*Team)
x := make(map[string][]*Room)
toms:= make(map[string]*managedTeam)
managed := make(map[string]map[string]*managedTeam)
roomteams := []*Team{}
rams := []*Room{}
for rows.Next() {
    room := &Room{}
    team := &Team{}
    err = rows.Scan(
      &room.Id,
      &room.Teamid,
      &room.Displayname,  
      &team.Id,
      &team.Displayname,
      
    )
    if err != nil {
		return nil, nil, nil, err
	}
    roomteams = append(roomteams, team)
    rams = append(rams, room)
    
    
    y[team.Name] = append(y[team.Name], team) 
    x[team.Name] = append(x[team.Name], room)
    toms["key"] = &managedTeam{teams: y, rooms: x}
    managed["mykeys"] = toms

  }
        return rams, roomteams, managed , nil

}


func (store *dbStore) GetRoomTest(codes []string) (*Room, error) {
    var r Room
    // this calls sql.Open, etc.
    //var aut Author
    err := store.db.QueryRow("SELECT * FROM rooms WHERE teamid = any($1)", pq.Array(codes)).Scan(&r.Id, &r.CreateAt, &r.Teamid, &r.Name, &r.Type, &r.Displayname, &r.Authorid)
    if err != nil {
        return &Room{}, err
    } else {
        return &r, nil
    }
}


func (store *dbStore) GetRoomsByTeamId(id string, email string) (*Room, error) {
    var r Room
    // this calls sql.Open, etc.
    //var aut Author

    // note the below syntax only works for postgres
    err := store.db.QueryRow("SELECT * FROM rooms WHERE teamid = $1", id).Scan(&r.Id, &r.CreateAt, &r.Teamid, &r.Name, &r.Type, &r.Displayname, &r.Authorid)
    if err != nil {
        return &Room{}, err
    } else {
        return &r, nil
    }
}


func (store *dbStore) GetSession(hash string) (*Session, error) {
    var s Session
    // this calls sql.Open, etc.
    //var aut Author

    // note the below syntax only works for postgres
    err := store.db.QueryRow("SELECT * FROM sessions WHERE hash = $1", hash).Scan(&s.ID, &s.Authori, &s.Hash)
    if err != nil {
        return &Session{}, err
    } else {
        return &s, nil
    }
}


func (store *dbStore) GetSessionByName(name string) (*Session, error) {
    var s Session
    // this calls sql.Open, etc.
    //var aut Author

    // note the below syntax only works for postgres
    err := store.db.QueryRow("SELECT * FROM sessions WHERE authori = $1", name).Scan(&s.ID, &s.Authori, &s.Hash)
    if err != nil {
        return &Session{}, err
    } else {
        return &s, nil
    }
}







func (store *dbStore) GetRoomByEmail(email string) (*Room, error) {
    var r Room
    // this calls sql.Open, etc.
    //var aut Author

    // note the below syntax only works for postgres
    err := store.db.QueryRow("SELECT * FROM rooms WHERE authorid = $1", email).Scan(&r.Id, &r.Name, &r.CreateAt, &r.Teamid, &r.Type, &r.Displayname, &r.Authorid)
    if err != nil {
        return &Room{}, err
    } else {
        return &r, nil
    }
}

func (store *dbStore) GetCountRooms() (int64 , error) {
    var count int64
    // this calls sql.Open, etc.
    //var aut Author
    err := store.db.QueryRow("SELECT COUNT(*) FROM rooms").Scan(&count)
    // note the below syntax only works for postgres
    //err := store.db.QueryRow("SELECT authorid FROM rooms").Scan(&r.Authorid)
    if err != nil {
        return count, err
    } else {
        return count, nil
    }
}




func (store *dbStore) GetRoomByAuthor() (*Room, error) {
    var r Room
    // this calls sql.Open, etc.
    //var aut Author

    // note the below syntax only works for postgres
    err := store.db.QueryRow("SELECT authorid FROM rooms").Scan(&r.Authorid)
    if err != nil {
        return &Room{}, err
    } else {
        return &r, nil
    }
}


var stora Stora

// The store variable is a package level variable that will be available for
// use throughout our application code
var store Store

func InitStores(sa Stora) {
     stora = sa
}

/*
We will need to call the InitStore method to initialize the store. This will
typically be done at the beginning of our application (in this case, when the server starts up)
This can also be used to set up the store as a mock, which we will be observing
later on
*/
func InitStore(s Store) {
	store = s
}
