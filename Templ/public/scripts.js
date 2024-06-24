function show(id) {
  document.getElementById(id).classList.remove('hidden');
}

function hide(id) {
  document.getElementById(id).classList.add('hidden');
}

function setPending(id) {
  document.getElementById(id).classList.add('pending');
}

function unsetPending(id) {
  document.getElementById(id).classList.remove('pending');
}

function setDisabled(id) {
  document.getElementById(id).disabled = true;
}

function unsetDisabled(id) {
  const target = document.getElementById(id)

  if (!target) {
    alert("Target not found")
  }

  target.disabled = false;
}

function setSessionCookie(value) {
  //cookie:session_id
  // document.cookie = "session_id=" +
}

function closeDialog(id) {
  for (let i = 0; i < document.getElementsByClassName('dialog').length; i++) {
    document.getElementsByClassName('dialog')[i].close()
  }
}