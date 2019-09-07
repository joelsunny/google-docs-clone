import React from 'react';
import logo from './logo.svg';
import './App.css';
import Quill from 'quill';

const base_url = "http://";

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

    log(delta, oldDelta, source) {
        const logstream = document.getElementById("c" + this.props.id);
        const url = base_url + 'delta';
        logstream.innerHTML = "delta : " + JSON.stringify(delta) + "<br>" + "old: " + JSON.stringify(oldDelta) + "<br>" + "source: " + source + "<br>" + "-------------------<br><br>" + logstream.innerHTML;
        fetch(url, {
            method: 'POST',
            body : {
                "id" : this.props.id,
                "delta" : delta["ops"]
            }
        });
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

function App() {
    return (
    <div className="App">
        <Editor id="1" />
        <Editor id="2" />
    </div>
    );
}

export default App;
