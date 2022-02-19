import { Component } from "react";


class ChatInput extends Component {
    render() {
        return(
            <div className="ChatInput">
                <input onKeyDown={this.props.send} placeholder="Press Enter to send"/>
            </div>
        )
    }
}

export default ChatInput;