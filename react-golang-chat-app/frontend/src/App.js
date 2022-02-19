import Header from "./components/Header/Header";
import ChatHistory from './components/ChatHistory/ChatHistory';
import { Component } from "react";
import { connect, sendMsg} from './api';
import './App.css';
import ChatInput from "./components/ChatInput/ChatInput";

class App extends Component {

  constructor(props) {
    super(props);
    this.state = {
      chatHistory: [],
    }
  }

  componentDidMount() {
    connect((msg) => {
      this.setState(prevState => ({
        chatHistory: [...this.state.chatHistory, msg]
      }))
    });
  }


  send(event) {
    if(event.keyCode === 13) {
      sendMsg(event.target.value);
      event.target.value = "";
    }
  }


  render() {
    return (
      <div className="App">
            <Header/>
            <ChatHistory chatHistory={this.state.chatHistory} />
            <ChatInput send={this.send}/>
      </div>
    );
  }


}


export default App;