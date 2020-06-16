import React, {Component} from 'react';
import {Card, Header, Form, Input, Icon} from "semantic-ui-react";

let endpoint = "http://127.0.0.1:8080/dashboard/";



class TODOList extends Component {
    constructor(props) {
        super(props);

        this.state = {
            task: "",
            items: []
        };
    }

    componentDidMount() {
        this.getTask()
    };

    onChange = event => {
        this.setState({
            [event.target.name]: event.target.value
        });
    }

    onSubmit = () => {
        let { task } = this.state;
        if (task) {
            fetch(
                endpoint, 
                {
                    method: "POST",
                    body: JSON.stringify({ "title": task }),
                    headers: {
                        "Content-Type": "application/json",
                    }
                }
            )
            .then (res => {
                this.getTask();
                this.setState({
                    task: ""
                });
                console.log(res);
            });
        }
    };

    getTask = () => {
        fetch(endpoint)
        .then(data => data.json())
        .then(res => {
            if (res) {
                this.setState({
                    items: res.map(item => {
                        let color = "yellow";

                        if (item.message) {
                            color = "green";
                        }

                        return (
                            <Card key={item.title} color={color} fluid>
                                <Card.Content>
                                    <Card.Header textAlign="left">
                                        <div style={{ wordWrap: "break-word" }}>{item.title}</div>
                                    </Card.Header>

                                    <Card.Meta textAlign="right">
                                        <Icon 
                                            name="check circle"
                                            color="green"
                                            onClick={() => this.updateTask(item.title)}
                                        />
                                        <span style={{ paddingRight: 10 }}>Done</span>
                                        <Icon 
                                            name="undo"
                                            color="yellow"
                                            onClick={() => this.undoTask(item.title)}
                                        />
                                        <span style={{ paddingRight: 10 }}>Undo</span>
                                        <Icon
                                            name="delete"
                                            color="red"
                                            onClick={() => this.deleteTask(item.title)}
                                        />
                                        <span style={{ paddingRight: 10 }}>Delete</span>
                                    </Card.Meta>
                                </Card.Content>
                            </Card>
                        );
                    })
                });
            } else {
                this.setState({
                    items: []
                });
            }
        });
    };

    deleteTask = title => {
        fetch(endpoint + title, {method: "DELETE"})
          .then(res => {
            console.log(res);
            this.getTask();
          });
      };
    
    render() {
    return (
        <div>
        <div className="row">
            <Header className="header" as="h2">
            TO DO LIST
            </Header>
        </div>
        <div className="row">
            <Form onSubmit={this.onSubmit}>
            <Input
                type="text"
                name="task"
                onChange={this.onChange}
                value={this.state.message}
                fluid
                placeholder="Create Task"
            />
            </Form>
        </div>
        <div className="row">
            <Card.Group>{this.state.items}</Card.Group>
        </div>
        </div>
    );
    }
}

export default TODOList;