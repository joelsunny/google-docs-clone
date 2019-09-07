import React from 'react';
import logo from './logo.svg';
import './App.css';
import Quill from 'quill';

var io = new WebSocket("ws://192.168.0.103:8080");

class Editor extends React.Component {

    constructor(props) {
        super(props);

        this.log = this.log.bind(this);
        
    }

    componentDidMount() {
        var container = document.getElementById('e' + this.props.id);
        var quill = new Quill(container, {
            modules:{
                toolbar:false
            },
            placeholder: 'Compose an epic...',
            theme: 'snow'  // or 'bubble'
          });

          // add a listener for text-change event
          quill.on('text-change', this.log);
    }

    async deltaPropogate(msg) {
        console.log('sending delta to server');
        io.send(msg);
    }

    log(delta, oldDelta, source) {
        const logstream = document.getElementById("c" + this.props.id);
        logstream.innerHTML = "delta : " + JSON.stringify(delta) + "<br>" + "old: " + JSON.stringify(oldDelta) + "<br>" + "source: " + source + "<br>" + "-------------------<br><br>" + logstream.innerHTML;
        this.deltaPropogate(JSON.stringify(delta["ops"]));
    }

    render() {
        return(
            <div id={this.props.id} className="editor">
                <div className="controls">
                    controls
                </div>
                <div id={"e" + this.props.id} className="text-area">
                </div>
                <div id={"c" + this.props.id} className="console">
                    
                </div>
            </div>
        )
    }
}

class ServerView extends Editor {
    constructor(props) {
        super(props);
    }

    componentDidMount() {

        io.onopen = () => {
                            this.log("url: " + io.url);
                            this.log("connection status: " + io.readyState);
                        };
        
        io.onmessage = (msg) => {
                this.log(msg.data);
                console.log(msg);
        };
    }

    log(msg) {
        const logstream = document.getElementById("c" + this.props.id);
        logstream.innerHTML = "got: " + msg + "<br>" + "<br>" + logstream.innerHTML;
    }
}

function App() {
    return (
    <div className="App">
        <Editor id="1" />
        <ServerView id="2" />
    </div>
    );
}

export default App;
