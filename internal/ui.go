package internal

import (
	"net/http"
)

var page = []byte(`
  <!DOCTYPE html>
  <html>
    <head>
      <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.14.2/graphiql.css" />
      <script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/3.0.0/fetch.js"></script>
      <script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js"></script>
      <script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js"></script>
      <script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.14.2/graphiql.js"></script>
    </head>
    <body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">

    <input id="fieldButton" type="button" value="Show Fields" onclick="hide()" style="background-color:#f7f7f7;height:30px;border:2px solid black;width:100%">

    <div align="center" id="fields" style="display:none;">
      <input id="userId" placeholder="User ID" style="width:30%;border-radius:2px;height:20px;">      
      <input id="tkpdsessionid" placeholder="tkpd-sessionid" style="width:30%;border-radius:2px;height:20px;">
      <input id="auth" placeholder="Accounts Authentication Token" style="width:30%;border-radius:2px;height:20px;"><br><br>

      <input id="x-device" placeholder="x-device" style="width:30%;border-radius:2px;height:20px;">
      <input id="Content-MD5" placeholder="Content-MD5" style="width:30%;border-radius:2px;height:20px;">
      <input id="Authorization" placeholder="Authorization" style="width:30%;border-radius:2px;height:20px;"><br><br>
      <input id="Content-Type" placeholder="Content-Type" style="width:30%;border-radius:2px;height:20px;">
      <input id="x-app-version" placeholder="X-App-Version" style="width:30%;border-radius:2px;height:20px;">
      <input id="x-tkpd-akamai" placeholder="x-tkpd-akamai" style="width:30%;border-radius:2px;height:20px;"><br><br>

      Beta Bypass Header: <input type="checkbox" id="x-beta-bypass" name="Bypass Beta">
      <label for="x-beta-bypass"> (Tkpd-Pragma:tkpd-x-bypass-beta) </label><br><br>

      HMAC at GQL: <input type="checkbox" id="x-return-hmac-md5" name="x-return-hmac-md5">
      <label for="x-return-hmac-md5"> (Return HMAC and MD5 at GQL?)</label><br><br>
      
      <div style="border:1px dotted black;width:30%;text-align:left;backgournd-color:#adebeb">For Example:<br>User ID:3045010<br>Authentication Token: zEaDspxLSNCTa3wf1wnnxg 
      <br>tkpd-sessionid: f3LAWTH2DDs:APA91bEvA4S1HcTJso_WoJC <br>
      x-device: android-2.20<br>
      </div>
    </div>
      <div id="graphiql" style="height: 95vh;">Loading...</div>
      <script>
      function hide() {
        var x = document.getElementById("fields");
        if (x.style.display === "none") {
            x.style.display = "block";
            document.getElementById("fieldButton").value = 'Hide Fields'
        } else {
            x.style.display = "none";
            document.getElementById("fieldButton").value = 'Show Fields'
        }
    }
        function graphQLFetcher(graphQLParams) {
          const headerObject = {
            'origin':'m.tokopedia.com',
          }

          if (document.getElementById("Authorization").value) {
            headerObject['Authorization'] = document.getElementById("Authorization").value
          }

          if (document.getElementById("Content-MD5").value) {
            headerObject['Content-MD5'] = document.getElementById("Content-MD5").value
          }

          if (document.getElementById("Content-Type").value) {
            headerObject['Content-Type'] = document.getElementById("Content-Type").value
          } else {
            headerObject['Content-Type'] = 'application/json'
          }

          if (document.getElementById("x-return-hmac-md5").checked == true) {
            headerObject['x-return-hmac-md5'] = "x-return-hmac-md5"
          }

          if (document.getElementById("userId").value) {
            headerObject['Tkpd-UserId'] = document.getElementById("userId").value
          }

          if (document.getElementById("auth").value) {
            headerObject['Accounts-Authorization'] = 'Bearer ' + document.getElementById("auth").value
          }

          if (document.getElementById("tkpdsessionid").value) {
            headerObject['tkpd-sessionid'] = document.getElementById("tkpdsessionid").value
          }

          if (document.getElementById("x-device").value) {
            headerObject['x-device'] = document.getElementById("x-device").value
          }

          if (document.getElementById("x-app-version").value) {
            headerObject['x-app-version'] = document.getElementById("x-app-version").value
          }

          if (document.getElementById("x-tkpd-akamai").value) {
            headerObject['x-tkpd-akamai'] = document.getElementById("x-tkpd-akamai").value
          }

          if (document.getElementById("x-beta-bypass").checked == true) {
            headerObject['Tkpd-Pragma'] = "tkpd-x-bypass-beta"
          }
          
          return fetch("/graphql", {
            headers: headerObject,
            method: "post",
            body: JSON.stringify(graphQLParams),
            credentials: "include",
          }).then(function (response) {
            return response.text();
          }).then(function (responseBody) {
            try {
              return JSON.parse(responseBody);
            } catch (error) {
              return responseBody;
            }
          });
        }
        ReactDOM.render(
          React.createElement(GraphiQL, {fetcher: graphQLFetcher}),
          document.getElementById("graphiql")

        );
      </script>
    </body>
  </html>
`)

// GraphiQLHandler to handle ui
func GraphiQLHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(page)
}
