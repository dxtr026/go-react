var props = { foo: 'default' };
var Component = React.createClass({
  render : function(){
    return (
        <p>
        Rahul on is awesome {this.props.foo}
        </p>
      )
  }

});

var component = <Component {...props} foo={'override'} />;

// var component = React.createElement(Component, React.__spread({}, props, {foo: 'override'}))
React.render(component, document.getElementById('example'))
console.log(component.props.foo); // 'override'