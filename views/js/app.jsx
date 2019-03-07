//import SwaggerUI from "swagger-ui-react"
//import "swagger-ui-react/swagger-ui.css"

//export default App = () => <SwaggerUI url="https://petstore.swagger.io/v2/swagger.json" />

class App extends React.Component {
    render() {
        if (this.loggedIn) {
            return (<LoggedIn />);
        } else {
            return (<Home />);
        }
    }
}

class Home extends React.Component {
    render() {
        return (
            <div className="container">
                <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
                    <h1>Smart Kitchen</h1>
                    <p>Your smart home controls go here....</p>
                    <p>Sign in to edit control setup.</p>
                    <a onClick={this.authenticate} className="btn btn-primary btn-lg btn-login btn-block">Sign In</a>
                </div>
            </div>
        )
    }
}

class LoggedIn extends React.Component {
    constructor(props) {
        super(props);
    }
    render() {
        return (
            <div className="container">
                <div className="col-lg-12">
                    <br />
                    <span className="pull-right"><a onClick={this.logout}>Log out</a></span>
                    <h2>Edit your controls</h2>
                    <p>Let's edit your controls</p>

                </div>
            </div>
        )
    }
}

ReactDOM.render(<App />, document.getElementById('app'));