


//fetch('/api/v4/channels/1/posts')
  //.then((resp) => resp.json()) // Transform the data into json
  //.then(function(data) {
     //console.log(data)
    // Create and append the li's to the ul
    //})

//console.log(document.cookie)


$(function(){
   console.log(window.location.host)



        var msgBox = $("#tb_message");
        var msgBoxi = $("#species");

class EventEmitter {
  listeners = {}
  
  addListener(eventName, fn) {
    this.listeners[eventName] = this.listeners[eventName] || [];
    this.listeners[eventName].push(fn);
    return this;
  }
  // Attach event listener
  on(eventName, fn) {
    return this.addListener(eventName, fn);
  }

  // Attach event handler only once. Automatically removed.
  once(eventName, fn) {
    this.listeners[eventName] = this.listeners[eventName] || [];
    const onceWrapper = () => {
      fn();
      this.off(eventName, onceWrapper);
    }
    this.listeners[eventName].push(onceWrapper);
    return this;
  }

  // Alias for removeListener
  off(eventName, fn) {
    return this.removeListener(eventName, fn);
  }

  removeListener (eventName, fn) {
    let lis = this.listeners[eventName];
    if (!lis) return this;
    for(let i = lis.length; i > 0; i--) {
      if (lis[i] === fn) {
        lis.splice(i,1);
        break;
      }
    }
    return this;
  }

  // Fire the event
  emit(eventName, ...args) {
    let fns = this.listeners[eventName];
    if (!fns) return false;
    fns.forEach((f) => {
      f(...args);
    });
    return true;
  }

  listenerCount(eventName) {
    let fns = this.listeners[eventName] || [];
    return fns.length;
  }

  // Get raw listeners
  // If the once() event has been fired, then that will not be part of
  // the return array
  rawListeners(eventName) {
    return this.listeners[eventName];
  }

}

const CHANGE_EVENT = 'change';

class PostStoreClass extends EventEmitter {
      constructor() {
        super();
        this.selectedPostId = null;
        this.postsInfo = {};
        this.latestPageTime = {};
        this.earliestPostFromPage = {};
        this.currentFocusedPostId = null;
    }
    emitChange() {
        this.emit(CHANGE_EVENT);
    }

}




      class MySocket {
  constructor(path) {
 
    this.path = path; 

    this.listeners = []

    this.emitter = new EventEmitter()

    this.eventListener = {};

    this.on = (event, cb) => this.eventListener[event] = cb;
    
    this.noupath = this.path.substring(0, this.path.length - 1);

    this.serviceLocation = "wss://" + window.location.host + this.noupath;

      try {
      this.websocket = new WebSocket(this.serviceLocation);
      
      this.socket = this.websocket;

        } catch (e) {
            console.log(e);
        }


    

    //this.socket.onclose = this.onClose.bind(this);
    this.socket.onerror = this.onError.bind(this);
    this.socket.onopen = this.onOpen.bind(this);
    this.socket.onmessage = this.onMessage.bind(this);

    this.emit = (event, data) => {
        let rawData = JSON.stringify(event, data);

        this.socket.send(rawData);
    }

    





  } 

onOpen(evt){
console.log('connection etablie');

    


}


onError(evt){
      console.log("Socket error: ");

     

  }

  onClose(evt){
    

window.alert('Websocket connection closed.');
     

  }


  

  onMessage(evt){

  try {









            let data = JSON.parse(evt.data);


            if (data) {
                
                let cb = this.eventListener[data.event];
                if (cb)
                    cb(data);
            }


            
        } catch (e) {
            console.log(e);
        }












  } 






  

}

Element.prototype.remove = function() {
    this.parentElement.removeChild(this);
}
NodeList.prototype.remove = HTMLCollection.prototype.remove = function() {
    for(var i = this.length - 1; i >= 0; i--) {
        if(this[i] && this[i].parentElement) {
            this[i].parentElement.removeChild(this[i]);
        }
    }
}

 var conn = new MySocket(window.location.pathname);

var PostStore = new PostStoreClass();







//console.log(conn);
var sample = new Array();
var racineOfRoom = [];
var lastMsgUsername;
var myUsername;
var currentUser;
var currentRoom;
var div_responses = document.getElementById('div_responses');
var div_emojis = document.getElementById('div_emojis');
var div_posts = document.getElementById('div_posts');
var lb_connected = document.getElementById('lb_connected_counter');
var ul_users_list = document.getElementById('ul_users_list');
var f_name = document.getElementById('f_name');
var div_login = document.getElementById('login');
var div_acc_create = document.getElementById('acc_create');
var f_input = document.getElementById('f_input');
var fi_input = document.getElementById('fi_input');
var fi_input_id = document.getElementById('fi_input_id');
var tb_name = document.getElementById('tb_name');
//div_login.style.display = "none";
var tb_message = document.getElementById('tb_message');

f_name.onsubmit = (e) => {
    e.preventDefault();
    
    myUsername = tb_name.value;
    conn.emit({
        event: 'login',
        data: {
            username: myUsername,
            password: tb_password.value,
            cookie: document.cookies
        },
    });
};


fi_input.onsubmit = (e) => {
    e.preventDefault();
    
    myUsername = fi_input_id.value;
    conn.emit({
        event: 'suscribe',
        data: {
            username: myUsername
            
        },
    });
};




tb_name.oninput = (e) => {
    let cvalue = e.target.value;
    setTimeout(() => {
       
        if (tb_name.value == cvalue) {
            conn.emit({
                event: 'checkUsername',
                data: cvalue
            });
        }
    }, 300);
};





function getTime(timestamp) {
    function btf(inp) {
        if (inp < 10)
            return '0' + inp;
        return inp;
    }
    var date = new Date(timestamp * 1000),
        y = date.getFullYear(),
        m = btf(date.getMonth() + 1),
        d = btf(date.getDate()),
        h = btf(date.getHours()),
        min = btf(date.getMinutes()),
        s = btf(date.getSeconds());
    return `${d}.${m}.${y} - ${h}:${min}:${s}`;
}


function updateUsersList(usersMap) {

    ul_users_list.innerHTML = '';
    Object.keys(usersMap).forEach(uname => {
         console.log(usersMap[uname].Color)
    let elem = document.createElement('li');
       
        elem.innerText = uname;
        elem.style.color = usersMap[uname].Color; 
        //ul_users_list.style.color = "orangered";
        ul_users_list.appendChild(elem);

        //ul_users_list.appendChild(elem);
    });
}




function appendPosts(content) {


var converter = new showdown.Converter({headerLevelStart: 3, strikethrough: true, emoji: true, underline: true,});




         let div = document.createElement('div');
         div.className = 'message_tile';
         div.id = 'container' + content.id;

         let divTitle = document.createElement('div');
         divTitle.className = 'head';

    if (lastMsgUsername != content.userid) {
        let uname = document.createElement('div');
        uname.innerText = content.userid;
        uname.className = 'username';
        uname.style.color = content.color;
        divTitle.appendChild(uname);

        let time = document.createElement('div');
        time.innerText = getTime(content.createat);
        time.className = 'time';
        divTitle.appendChild(time);
        let tim = document.createElement('div');
        tim.innerText = currentRoom.Displayname;
        tim.className = 'room';
        tim.style.color = "red";
        tim.style.marginLeft = "40px";
        tim.style.marginTop = "10px"; 
        divTitle.appendChild(tim);

        div.appendChild(divTitle);
    }

    let message = document.createElement('div');
    message.id = 'message' + content.id;
    if(content.markdown === undefined){
    let converted = converter.makeHtml(content.message);
    message.innerHTML = content.message;
    }else{
    let converted = converter.makeHtml(content.markdown); 
    message.innerHTML = content.message;
    } 
    message.className = 'message';
    

    div.appendChild(message);

    if (content.userid == myUsername) {
        let messageActionDiv = document.createElement('div');
        let messageActionDivEdit = document.createElement('div');
        let messageEdit = document.createElement('div');
        messageEdit.style.display = "none";
        let deleteLink = document.createElement('a');
        let editLink = document.createElement('a');
        let edit = document.createElement('div');
        let f = document.createElement("form");
        f.setAttribute('id', 'editid') 
 
        let inpi = document.createElement("input");
        inpi.setAttribute('type',"text");
        inpi.setAttribute('id',"inputedit" + content.id);
        inpi.setAttribute('value', content.message);
        inpi.style.width = "580px";
        let si = document.createElement("input"); //input element, Submit button
        si.setAttribute('type',"button");
        si.setAttribute('value',"Submit"); 
        si.setAttribute('id',"editbutid");
   
        deleteLink.onclick = () => {
            conn.emit({
                event: 'deleteMessage',
                data: { 
                      msgid: content.id
         
                }
            });
        };
        editLink.onclick = () => {
            messageEdit.className = "edit";
            messageEdit.id = "editor";
            messageEdit.style.display = "block";
            f.appendChild(inpi);
            f.appendChild(si);
            //console.log(inpi)
            //inpi.setAttribute('value', content.message);
            edit.appendChild(f);
            messageEdit.appendChild(edit);
            div.appendChild(messageEdit); 
            
    let mytarget = document.getElementById('inputedit'+ content.id);


if(mytarget !== null) {
mytarget.addEventListener("change", myFunction);



};


function myFunction() {

conn.emit({
                event: 'editMessage',
                data: { 
                    msgid: content.id,
                    mycontent: mytarget.value 
                }
            });
let natarget = document.getElementById('editbutid');
natarget.onclick = () => { 
message.innerHTML = mytarget.value;
document.getElementById('editor').style.display = "none";
}

}



}
            

            




        editLink.innerText = "edit"; 
        deleteLink.innerText = "remove";
        messageActionDivEdit.className = "editLink";
        messageActionDiv.className = "deleteLink";
        messageActionDiv.appendChild(deleteLink);
        messageActionDivEdit.appendChild(editLink);
        div.appendChild(messageActionDiv);
        div.appendChild(messageActionDivEdit);
    }

    div_responses.appendChild(div);
    div.scrollIntoView(); 
    lastMsgUsername = content.userid;






}




function appendMessage(msgEvent) {
//console.log(msgEvent)
var converter = new showdown.Converter({headerLevelStart: 3, strikethrough: true, emoji: true, underline: true,});
         let div = document.createElement('div');
         div.className = 'message_tile';
         div.id = 'container' + msgEvent.Id;

         let divTitle = document.createElement('div');
         divTitle.className = 'head';

    if (lastMsgUsername != msgEvent.author) {
        let uname = document.createElement('div');
        uname.innerText = msgEvent.author;
        uname.className = 'username';
        uname.style.color = msgEvent.Color;
        divTitle.appendChild(uname);

        let time = document.createElement('div');
        time.innerText = getTime(msgEvent.timestamp);
        time.className = 'time';
        divTitle.appendChild(time);
        let tim = document.createElement('div');
        tim.innerText = msgEvent.room;
        tim.className = 'room';
        tim.style.color = "red";
        tim.style.marginLeft = "40px";
        tim.style.marginTop = "10px"; 
        divTitle.appendChild(tim);

        div.appendChild(divTitle);
    }

    let message = document.createElement('div');
    message.id = 'message' + msgEvent.Id;
    let converted = converter.makeHtml(msgEvent.content);
    message.innerHTML = msgEvent.content;
    message.className = 'message';
    if (msgEvent.content.includes('@' + myUsername)) {
        message.className += ' highlighted';
        let ops = {
            body: msgEvent.content,
            icon: "https://camo.githubusercontent.com/6490d7a3b892fd102a9ef1718aec3c41838639ee/68747470733a2f2f7a656b726f2e64652f7372632f676f5f636861745f6c6f676f2e706e67"
        }
        if (!("Notification" in window)) {
            alert("This browser does not support desktop notification");
        } else if (Notification.permission === "granted") {
            var notification = new Notification("You got pinged", ops);
        } else if (Notification.permission !== "denied") {
            Notification.requestPermission(function (permission) {
                if (permission === "granted") {
                    var notification = new Notification("You got pinged", ops);
                }
            });
        }
    }

    div.appendChild(message);

    if (msgEvent.author == myUsername) {
        let messageActionDiv = document.createElement('div');
        let deleteLink = document.createElement('a');
        deleteLink.onclick = () => {
            conn.emit({
                event: 'deleteMessage',
                data: { 
                    msgid: msgEvent.Id
                }
            });
        };
        deleteLink.innerText = "remove";
        messageActionDiv.className = "deleteLink";
        messageActionDiv.appendChild(deleteLink);
        div.appendChild(messageActionDiv);
    }

    div_responses.appendChild(div);
    div.scrollIntoView(); 
    lastMsgUsername = msgEvent.author;

}

conn.on('usernameState', (data) => {
    console.log(data)

});



conn.on('createCookie', (data) => {

document.cookie = data.data;






});


conn.on('message', (data) => {
console.log(data)
appendMessage(data.data);



  
});




conn.on('messageDeleted', (data) => {

let msg = $('#message' + data.data.id);
let container = $('#container' + data.data.id);

msg.innerHTML = '<p class="status_msg">deleted</p>';
console.log(msg);
document.getElementsByClassName('deleteLink')[0].remove();

});



function estCerises(fruit) {
  return  fruit.Teamid === '1597969999';
}

function estC(fruit) {
  return  fruit.Id === '1597969999';
}

var groupBy = function(xs, key) {
  return xs.reduce(function(rv, x) {
    (rv[x[key]] = rv[x[key]] || []).push(x);
    return rv;
  }, {});
};

function remove_duplicates(arr) {
    var obj = {};
    var ret_arr = [];
    for (var i = 0; i < arr.length; i++) {
        obj[arr[i]] = true;
    }
    for (var key in obj) {
        ret_arr.push(key);
    }
    return ret_arr;
}

let modalios = null;

conn.on('clientConnect', (data) => {

//let divin = document.createElement('div');
const tierget = document.querySelector('#b1');


//console.log(evt.target.value);
const closeEmojis = function (event){
if (modalios === null) return 
event.preventDefault()

data.data.emo.forEach(uname => {

let spaniol = document.getElementById(uname.key); 

if(spaniol !== null){
div_emojis.removeChild(spaniol);
spaniol.innerHTML = "";
}
div_emojis.style.display = "none";
div_emojis.innerHTML = "";
});


modalios.removeEventListener('click', closeEmojis)

modalios = null

}






const openEmojis = function (event){
let index = 0;
data.data.emo.forEach(uname => {
      







let spaniol =  document.createElement('span');
spaniol.style.marginLeft = "20px";
spaniol.innerHTML = uname.valeur;
spaniol.id = uname.keyi;
div_emojis.style.display = "block";
div_emojis.appendChild(spaniol);

index++;
//div_emojis.removeChild(spaniol);
});

modalios = tierget
modalios.addEventListener('click', closeEmojis)


}


tierget.addEventListener("click", openEmojis )



currentUser = data.data.author.username
const tirget = document.querySelector('#cetuser');
tirget.innerText = currentUser;
currentRoom = data.data.palis
console.log("currentroom", data.data.palis)
//var urloPath = window.location.pathname.split("/");
//var fo_part = urloPath[2];

//console.log(groupBy(data.data.roomteams, 'Id'));

if(data.data.tasto !== undefined){

const torget = document.querySelector('#palou');
Object.values(data.data.tasto).forEach(uname => {
uname.forEach(el => {
console.log(el)
let li = document.createElement('li');
let newlink = document.createElement('a');
let span =  document.createElement('span');
span.innerText = "Switch to";
newlink.setAttribute("href", el);
newlink.innerText = el;
newlink.appendChild(span);
li.appendChild(newlink);
torget.appendChild(li)

});

});
//console.log(groupBy(data.data.roomiso, 'Teamid'));
}

if(data.data.test !== undefined){




var mapi = new Map(Object.entries(data.data.test));
//console.log(mapi.get('serveur'));
//console.log(mapi.get('tamaris'));
//console.log(mapi.get('tyrolien'));

for (let [key, value] of Object.entries(data.data.test)) {
  //console.log(key, value)
  

var option1 = document.createElement("option");
option1.text = key;
option1.value = key;
option1.className = "boldoptioni";



var select = document.getElementById("chatroomi");
select.appendChild(option1);





}
Object.values(data.data.test).forEach(uname => {

uname.forEach(el => {
//console.log(el)
var option2 = document.createElement("option");
option2.text = el.Displayname;
option2.value = el.Displayname;
option2.className = "boldoptiona";
var select = document.getElementById("chatrami");
select.appendChild(option2);

});

});



}
if(data.data.roomteams !== undefined){

 //console.log(groupBy(data.data.tosts, 'Teamid')); 
//data.data.roomteams.rooms = [];
//data.data.roomteams.rooms
//for (let [key, value] of Object.entries(groupBy(data.data.roomteams, 'Id'))) {
  
         //data.data.roomteams.rooms.push(value)  
         //cube.rooms = value;
         //console.log(value);
      

//}//
//console.log("hhhhhh", data.data.roomteams.rooms); 
//console.log(data.data.roomteams.rooms.flat()); 


data.data.roomteams.forEach(uname => {
//console.log("serv", uname)







});



var mestests = document.getElementsByClassName("boldoptioni");


Array.prototype.forEach.call(mestests, function(el) {
    // Do stuff here

    el.addEventListener("click", function(evt) { 

var room = $('#chatroomi option:selected').val();
console.log(room);


}, false);
   


});


}





if(data.data.tosts !== undefined){


data.data.tosts.forEach(uname => {
//console.log("rooms", uname.Displayname)



});


}
if(data.data.palis !== undefined){

//console.log(data.data.palis)

}
if(data.data.pals !== undefined){
//data.data.pals['serveur'] = data.data.roomteams[0];
//data.data.pals['tamaris'] = data.data.roomteams[1];
//data.data.pals['tyrolien'] = data.data.roomteams[2];
//data.data.pals['serveur'].room = data.data.tosts[0];
//data.data.pals['tamaris'].room = data.data.tosts[1];
//data.data.pals['tyrolien'].room = data.data.tosts[2];
console.log("curentserver", data.data.pals)

}

if (data.data.history) {
        data.data.history.forEach(msg => appendMessage(msg.data));
    }
if (data.data.hposts) {
console.log(data.data.hposts);    
PostStore.postsInfo = Object.values(groupBy(data.data.hposts, 'parentid'))
//Object.values(groupBy(data.data.hposts, 'parentid')).forEach(uname => {






data.data.hposts.forEach(msg => appendPosts(msg));

PostStore.postsInfo.forEach(uname => {

uname.forEach(el => {



//console.log(el.hasOwnProperty('id'))
});



});




        //data.data.hposts.forEach(msg => appendPosts(msg));
    }
    currentRoom = data.data.author.Displayname
    currentUser = data.data.author.username;
    myUsername = data.data.author.username;
    div_login.style.top = `-${window.innerHeight}px`;
    setTimeout(() => {
        div_login.style.display = "none";
    }, 750);
    tb_name.autofocus = false;
    tb_message.autofocus = true;
    tb_message.focus();
    lb_connected.innerText = data.data.nclients;











});

conn.on('connected', (data) => {
console.log(data)
let elem = document.createElement('p');
    elem.innerText = `[${data.data.author.username} CONNECTED]`;
    elem.className = 'message_tile status_msg';
    div_responses.appendChild(elem);
    updateUsersList(data.data.clients);






currentRoom = data.data.room;



});



conn.on('disconnected', (data) => {
console.log(data);


});

conn.on('subscribe', (data) => {
console.log(data);


});




var test = document.getElementById("test");

$("#logout").click(function(e){
        //This will return true after the first click 
        //and preventDefault won't be called.
        if(!$(this).hasClass('nav_returnlink'))
            e.preventDefault();
console.log(test.innerHTML);
conn.emit({
        event: 'disconnected',
        data: test.innerHTML,
    });        
        $(this).addClass('nav_returnlink');
        $(this).unbind('click');
    });

$("#f_input").submit(function(e){

          e.preventDefault();
          if (!msgBox.val()) return false;
          if (!conn.websocket) {
            alert("Error: There is no socket connection.");
            return false;
          }
          conn.emit({
        event: 'message',
        data: msgBox.val(),
    });
          //conn.websocket.send("hhhhh");
          msgBox.val("");
          return false;

});


var test = document.getElementById("menubutton");
var testu = document.getElementById("fi_input"); 
var testuval = document.getElementById("fi_input_id");

var test22 = document.getElementById("menubutton22");
var testu22 = document.getElementById("fi_input22"); 
var testuval22 = document.getElementById("fi_input_id22");


test.addEventListener("click", function(evt) {
evt.preventDefault();
console.log(evt);
testu.style.display = "block"; 


}, false);

testu.addEventListener("submit", function(evt) {
evt.preventDefault();
console.log(testuval.value);

       conn.emit({
        event: 'subscribe',
        data: testuval.value,
    });


testuval.value = "";
}, false);

test22.addEventListener("click", function(evt) {
evt.preventDefault();

testu22.style.display = "block"; 


}, false);

testu22.addEventListener("submit", function(evt) {
evt.preventDefault();
console.log(testuval22.value);

          conn.emit({
        event: 'public',
        data: testuval22.value,
    });   



testuval22.value = "";
}, false);

let modal = null;
let modali = null;

const openModaliii = function (event){
event.preventDefault()
const paloi = document.querySelector('#paloi');
const targeti = document.querySelector('ul[role="menu"]');
targeti.style.display = "block";
targeti.setAttribute('aria-expanded', true);
modali = paloi
modali.addEventListener('click', closeModaliii)

}


document.querySelectorAll('.dropdown-toggle').forEach(a => {
  //alert(`${item} is at index ${index} in ${array}`);


   a.addEventListener('click', openModaliii)
   

});

const closeModaliii = function (event){
if (modali === null) return 
event.preventDefault()
const targeti = document.querySelector('ul[role="menu"]');
targeti.style.display = "none";
targeti.setAttribute('aria-expanded', false);

modali.removeEventListener('click', closeModaliii)

modali = null

}

const openModal = function (event){
event.preventDefault()
///console.log(event.target.getAttribute('href'))
const targeti = document.querySelector(event.target.getAttribute('href'));
targeti.style.display = null;
targeti.removeAttribute('aria-hidden');
targeti.setAttribute('aria-modal', true);
modal = targeti;
modal.addEventListener('click', closeModal)
modal.querySelector('.js-close-modal').addEventListener('click', closeModal)
modal.querySelector('.js-modal-stop').addEventListener('click', stopPropagation)
}


document.querySelectorAll('.js-modal').forEach(a => {
  //alert(`${item} is at index ${index} in ${array}`);
   a.addEventListener('click', openModal)
});


const closeModal = function (event){
if (modal === null) return 
event.preventDefault()
modal.style.display = "none";
modal.setAttribute('aria-hidden', true);
modal.removeAttribute('aria-modal');
modal.removeEventListener('click', closeModal)
modal.querySelector('.js-close-modal').removeEventListener('click', closeModal)
modal.querySelector('.js-modal-stop').removeEventListener('click', stopPropagation)
modal = null

}

const stopPropagation = function (event){
event.stopPropagation()
}




      }); 








window.addEventListener("load", () => {
    console.log('yyy')

    // Make sure name is focused on start
    //$("#name").focus();
}); 

