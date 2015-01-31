var props = { foo: 'default' };
var Component = React.createClass({displayName: "Component",
  render : function(){
    return (
        React.createElement("p", null, 
        "Rahul ", this.props.foo
        )
      )
  }

});

var component = React.createElement(Component, React.__spread({},  props, {foo: 'override'}));

// var component = React.createElement(Component, React.__spread({}, props, {foo: 'override'}))
React.render(component, document.getElementById('example'))
console.log(component.props.foo); // 'override'