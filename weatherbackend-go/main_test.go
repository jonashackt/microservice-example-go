package main

import (
	"bytes"
	"encoding/json"
	"github.com/jonashackt/microservice-example-go/weatherbackend-go/domain"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var app App

func TestMain(m *testing.M) {
	app = App{}
	app.Initialize()

	code := m.Run()

	os.Exit(code)
}

func TestGetSenseInThat(t *testing.T) {
	// Given
	name := "Jonas"
	request, _ := http.NewRequest("GET", "/"+name, nil)

	// When
	response := executeRequest(request)

	// Then
	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "Hello "+name+"! This is a RESTful HttpService written in Go. Try to use some other HTTP verbs (don´t say 'methods' :P )\n" {
		t.Errorf("Expected a specific string. Got %s", body)
	}
}

func TestGeneralOutlook(t *testing.T) {
	// Given
	weather := domain.Weather{"99423", "blue"}
	weatherBuffer := new(bytes.Buffer)
	json.NewEncoder(weatherBuffer).Encode(weather)
	// bytes.NewBuffer(weather)

	request, _ := http.NewRequest("POST", "/weather/general/outlook", weatherBuffer)

	// When
	response := executeRequest(request)

	// Then
	checkResponseCode(t, http.StatusCreated, response.Code)

	//assert.Empty(t, c.Errors)
	generalOutlookExpected := domain.GeneralOutlook{"Weimar", "Germany", "BestStationInTown"}
	assert.Contains(t, response.Header().Get("Content-Type"), "application/json; charset=utf-8")

	var generalOutlookActual domain.GeneralOutlook
	json.NewDecoder(response.Body).Decode(generalOutlookActual)
	assert.Equal(t, generalOutlookExpected, generalOutlookActual)

}

/*@Test
public void testWithRestAssured() {

Weather weather = new Weather();
weather.setFlagColor("blue");
weather.setPostalCode("99425");
weather.setProduct(Product.ForecastBasic);
weather.addUser(new User(27, 4300, MethodOfPayment.Bitcoin));
weather.addUser(new User(45, 500300, MethodOfPayment.Paypal));
weather.addUser(new User(67, 60000300, MethodOfPayment.Paypal));

given() // can be ommited when GET only
.contentType(ContentType.JSON)
.body(weather)
.when() // can be ommited when GET only
.post("http://localhost:8080/weather/general/outlook")
.then()
.statusCode(HttpStatus.SC_OK)
.contentType(ContentType.JSON)
.assertThat()
.equals(IncredibleLogic.generateGeneralOutlook());

public static GeneralOutlook generateGeneralOutlook() {
        GeneralOutlook generalOutlook = new GeneralOutlook();
        generalOutlook.setCity("Weimar");
        generalOutlook.setDate(Date.from(Instant.now()));
        generalOutlook.setState("Germany");
        generalOutlook.setWeatherStation("BestStationInTown");
        return generalOutlook;
    }

GeneralOutlook outlook = given() // can be ommited when GET only
.contentType(ContentType.JSON)
.body(weather).post("http://localhost:8080/weather/general/outlook").as(GeneralOutlook.class);

assertEquals("Weimar", outlook.getCity());
}*/

func executeRequest(request *http.Request) *httptest.ResponseRecorder {
	responseRecorder := httptest.NewRecorder()
	app.Router.ServeHTTP(responseRecorder, request)
	return responseRecorder
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
