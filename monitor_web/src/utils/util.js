export function open_page(url) {
  let div = document.createElement("a");
  div.href = url
  div.target = "_blank"
  div.id = "windowOpen"
  div.style = "display:none"
  document.getElementsByTagName("body")[0].appendChild(div);
  document.getElementById('windowOpen').click();
  document.getElementById("windowOpen").parentNode.removeChild(document.getElementById("windowOpen"));
}
