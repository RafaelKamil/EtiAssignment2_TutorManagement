//fucntion to populate tutor data from the  selected dropdown to text boxes
function show(ele) {
    // GET THE SELECTED VALUE FROM <select> ELEMENT AND SHOW IT.
    var email = document.getElementById('email');
    var name = document.getElementById('name')
    var descriptions = document.getElementById('descriptions')

    emails = JSON.parse(ele.value)
    listemails.innerHTML = 'Tutor Email Selected: <b>' + ele.options[ele.selectedIndex].text;
    ChosenEmail = ele.options[ele.selectedIndex].text
    email.value = emails.email
    name.value = emails.name
    descriptions.value = emails.descriptions
}


window.onload = populateSelect();
//fucntions to populate a data from database to the dropdown list
function populateSelect() {

    // CREATE AN XMLHttpRequest OBJECT, WITH GET METHOD.
    var xhr = new XMLHttpRequest(),
        method = 'GET',
        overrideMimeType = 'application/json',
        url = 'http://10.31.11.12:9181/api/v1/tutor/GetAllTutor';

    // ADD THE URL OF THE FILE.
    xhr.onreadystatechange = function() {
        if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            // PARSE JSON DATA.
            var listemails = JSON.parse(xhr.responseText);
            var ele = document.getElementById('sel');
            for (var i = 0; i < listemails.length; i++) {
                emailJsonString = JSON.stringify(listemails[i])
                    // BIND DATA TO <select> ELEMENT.
                ele.innerHTML = ele.innerHTML +
                    "<option value='" + emailJsonString + "''>" + listemails[i].email + '</option>';
            }
        }
    };
    xhr.open(method, url, true);
    xhr.send();
}
// This to validate Create form
function validateForm() {
    var Name = document.forms["createtutor"]["Name"].value;
    var Email = document.forms["createtutor"]["Email"].value;
    var Descriptions = document.forms["createtutor"]["Descriptions"].value;
    //error message if name text box is not filled out
    if (Name == "" || Name == null) {
        alert("Name must be filled out");
        return false;
        //error Email if name text box is not filled out
    } else if (Email == "" || Email == null) {
        alert("Email must be filled out");
        return false;
        //error Descriptions if name text box is not filled out
    } else if (Descriptions == "" || Descriptions == null) {
        alert("Descriptions must be filled out");
        return false;
    }

}
//this functions to update tutor account by email
var ChosenEmail
async function updateTutor() {
    //get element by update form id
    var form = document.getElementById('updateTutor')
        //listener whenever buttons type="submit" is click
    form.addEventListener('submit', function(e) {
        e.preventDefault()
        var name = document.getElementById('name').value
        var email = document.getElementById('email').value
        var descriptions = document.getElementById('descriptions').value
            //fetch the update a tutor by email api
        fetch("http://10.31.11.12:9181/api/v1/tutor/UpdateTutorAccountByEmail/" + ChosenEmail, {
                method: 'PUT',
                body: JSON.stringify({
                    name: name,
                    email: email,
                    descriptions: descriptions
                }),
                headers: {
                    "Content-Type": "application/json; charset=UTF-8"
                }
            })
            //this to shows error code
            .then(function(response) {
                if (response.status != 202) {
                    alert("Unable To Update Tutor Account!")
                    return
                } else {
                    alert("Successfuly Update Tutor Account!")
                }
            })
    })
}
updateTutor()
    //this to delete tutor account by email
async function deleteTutor() {
    //get element by deletetutor form id
    var form = document.getElementById('deletetutor')
        //listener whenever button is clicked
    form.addEventListener('click', function(e) {
        //prevent the page to refresh
        e.preventDefault()
            //get the value from the html page
        var name = document.getElementById('name').value
        var email = document.getElementById('email').value
        var descriptions = document.getElementById('descriptions').value
            //fetch the delete a tutor api
        fetch("http://10.31.11.12:9181/api/v1/tutor/DeleteTutorAccountByEamil/" + ChosenEmail, {
                method: 'DELETE',
                body: JSON.stringify({
                    name: name,
                    email: email,
                    descriptions: descriptions
                }),
                headers: {
                    "Content-Type": "application/json; charset=UTF-8"
                }
            })
            //this to shows error message
            .then(function(response) {
                if (response.status != 202) {
                    alert("Unable To Delete Tutor Account!")
                    return
                } else {
                    alert("Successfuly Delete Tutor Account!")
                        //calling 3.4 api 
                    fetch("http://10.31.11.12:9141/api/v1/deleteassignedtutor/" + ChosenEmail, {
                        method: 'DELETE',
                        body: JSON.stringify({
                            email: email,
                        }),
                        headers: {
                            "Content-Type": "application/json; charset=UTF-8"
                        }
                    })
                }
            })
    })
}
deleteTutor()
    //this to create new tutor fucntions
async function createTutor() {
    //get element by deletetutor form id
    var form = document.getElementById('createNewTutor')
        //listener whenever button is clicked
    form.addEventListener('submit', function(e) {
        //prevent the page to refresh
        e.preventDefault()
            //get the value from the html page
        var name = document.getElementById('name').value
        var email = document.getElementById('email').value
        var descriptions = document.getElementById('descriptions').value
            //fetch the create a tutor api
        fetch("http://10.31.11.12:9181/api/v1/tutor/CreateNewTutor", {
                method: 'POST',
                body: JSON.stringify({
                    name: name,
                    email: email,
                    descriptions: descriptions
                }),
                headers: {
                    "Content-Type": "application/json; charset=UTF-8"
                }
            })
            .then(function(response) {
                if (response.status != 202) {
                    alert("Unable To Create Tutor Account!")
                    return
                } else if (response.status == 226) {
                    alert("The email you enter is already registerd")
                } else
                    alert("Successfully Create Tutor Account!")
            })
    })
}
createTutor()


let tutors;
//this to display all tutor from  search fucntions
displayAPIData();
document.getElementById("search").onsubmit = SearchEmail

async function displayAPIData() {
    //get all tutor api
    const response = await fetch(
        "http://10.31.11.12:9181/api/v1/tutor/GetAllTutor"
    );
    tutors = await response.json();

    displayDataToTable(tutors)
}
//this is the search tutor by email fucntions 
async function SearchEmail(e) {
    e.preventDefault();
    const searchEmail = document.getElementById("searchBox").value;

    const tutorByEmail = tutors.filter((tutor) => {
        if (searchEmail == "") {
            return true
        } else if (tutor.email == searchEmail) {
            return true;
        } else {
            return false;
        }
    });

    displayDataToTable(tutorByEmail)
}
//this to display all Registered tutor to a table list
function displayDataToTable(data) {
    const tableData = data
        .map(function(value) {
            return `<tr>
            <td>${value.tutor_id}</td>
              <td>${value.name}</td>
              <td>${value.email}</td>
              <td>${value.descriptions}</td>
          </tr>
          <tr>
          </tr>
          `;



        })
        .join("");

    //set tableBody to new HTML code
    const tableBody = document.querySelector("#tableBody");
    tableBody.innerHTML = tableData;
}