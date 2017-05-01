
// -*- JavaScript -*-


class PersonItem extends React.Component {
  render() {
    return (
      <tr>
        <td> {this.props.id}    </td>
        <td> {this.props.first} </td>
        <td> {this.props.last}  </td>
      </tr>
    );
  }
}

class PeopleList extends React.Component {
  constructor(props) {
    super(props);
    this.state = { people: [] };
  }

  componentDidMount() {
    this.serverRequest =
      axios
        .get("/people")
        .then((result) => {
           this.setState({ people: result.data });
        });
  }

  render() {
    const people = this.state.people.map((person, i) => {
      return (
        <PersonItem key={i} id={person.Id} first={person.First} last={person.Last} />
      );
    });

    return (
      <div>
        <table><tbody>
          <tr><th>Id</th><th>First</th><th>Last</th></tr>
          {people}
        </tbody></table>

      </div>
    );
  }
}

ReactDOM.render( <PeopleList/>, document.querySelector("#root"));
