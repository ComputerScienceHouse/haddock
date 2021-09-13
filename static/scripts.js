function main() {
    const urlParams = new URLSearchParams(location.hash.substring(1));
    if (urlParams.has('length')) {
        const length = parseInt(urlParams.get("length"))
        if (length >= 16 && length <= 48) {
            document.getElementById("input_length").value = length
        }
    }

    onLengthSubmit();
    document.getElementById("input_length").addEventListener("keydown", function (event) {
        // Number 13 is the "Enter" key on the keyboard
        if (event.keyCode === 13) {
            // Cancel the default action, if needed
            event.preventDefault();
            // Trigger the button element with a click
            document.getElementById("genbutton").click();
        }
    });

}

function clipboardize(passwordID) {
      /* Get the text field */
        var element = document.getElementById(passwordID);
        var range, selection, worked;

          if (document.body.createTextRange) {
            range = document.body.createTextRange();
            range.moveToElementText(element);
            range.select();
          } else if (window.getSelection) {
            selection = window.getSelection();
            range = document.createRange();
            range.selectNodeContents(element);
            selection.removeAllRanges();
            selection.addRange(range);
          }

          try {
            document.execCommand('copy');
            console.log('text copied');
          }
          catch (err) {
            console.log('unable to copy text');
          }

}

function onLengthSubmit() {
        const length = document.getElementById("input_length").value;
        const xkcd = document.getElementById("input_xkcd").checked;
        location.hash = "#length="+length
        const url = xkcd ? '/api/v1/xkcd?length=' : '/api/v1/haddock?length='
        fetch(url + length)
            .then(response => response.json())
            .then(data => {
                document.getElementById("passwords").innerHTML = "";
                var passNum = 0;
                data.forEach(p => {
                    // Create the password item
                    var passwordText = document.createElement('li');
                    passwordText.setAttribute('id', 'password' + passNum);
                    passwordText.innerText = p
                    passwordText.setAttribute('onclick', 'clipboardize(\"password' + passNum + '\")');

                    // Append the newly constructed "node" to the password list
                    document.getElementById("passwords").appendChild(passwordText);

                    passNum++;
                })
            });
    }


