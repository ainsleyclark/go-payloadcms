/*
Package payloadfakes provides mock implementations of various Payload CMS services.
These mocks can be used in unit tests to simulate the behavior of the Payload CMS API
without making actual network requests.

## Services

The following mocks are provided:

  - MockCollectionService: A mock implementation of the CollectionService interface.
  - MockGlobalsService: A mock implementation of the GlobalsService interface.
  - MockService: A mock implementation of the Service interface.
  - MockMediaService: A mock implementation of the MediaService interface.

Each mock service allows you to define the behavior of its methods by setting the
corresponding function fields. By default, the methods return an empty payloadcms.Response
and a nil error.

## Usage

1. Create an instance of the mock service you need.
2. Define the behavior of the methods you are going to use by setting the corresponding function fields.
3. Use the mock service instance in your tests.

This allows you to fully control and assert the behavior of your code when interacting with the Payload CMS API.

## Example

func TestPayload(t *testing.T) {

	// Create a new mock collection service
	mockCollectionService := payloadfakes.NewMockCollectionService()

	// Define the behavior of the FindByID method
	mockCollectionService.FindByIDFunc = func(ctx context.Context, collection payloadcms.Collection, id int, out any) (payloadcms.Response, error) {
	    // Custom logic for the mock implementation
	    return payloadcms.Response{}, nil
	}

	// Use the mock collection service in your tests
	myFunctionUsingCollectionService(mockCollectionService)

}
*/
package payloadfakes
