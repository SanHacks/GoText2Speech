# GoText2Speech

### Simply :
Connect to API,
- `collectResponse()`

**Create Struct,**
  `var quotes []string`
-

**-Map Struct Unmarshal JSON**

`json.Unmarshal(body, &quotes)`
      
**[Pass Struct to `Speak()` Function]()**
`speech := htgotts.Speech{Folder: "audio", Language: voices.English, Handler: &handlers.Native{}}`
Here The Example Is a **"RON SWANSON"** Sarcastic Bot

