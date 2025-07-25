package sendtables

// Property is an auto-generated interface for property, intended to be used when mockability is needed.
// property wraps a flattenedPropEntry and allows registering handlers
// that can be triggered on a update of the property.
type Property interface {
	// Name returns the property's name.
	Name() string
	// Value returns the current value of the property.
	Value() PropertyValue
	// OnUpdate registers a handler for updates of the property's value.
	//
	// The handler will be called with the current value upon registration.
	OnUpdate(handler PropertyUpdateHandler)
	/*
	   Bind binds a property's value to a pointer.

	   Example:

	   	var i int
	   	property.Bind(&i, ValTypeInt)

	   This will bind the property's value to i so every time it's updated i is updated as well.

	   The valueType indicates which field of the PropertyValue to use for the binding.
	*/
	Bind(variable any, valueType PropertyValueType)
}
