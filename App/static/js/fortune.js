let addArgs = (i,e) => {
	let bb = document.getElementById("args");
	let rr = `<div id="delete-`+i+`"><input type="number" id=percent-"`+i+`" name="percent-`+i+`" min="0" max="100"><select name="type-` + i + `"><option value=""></option>`		
	e.forEach((element) => {
		rr += `<option value="`+element+`">`+element+`</option>`
	} ) 
	rr +=  `</select><p id="del-`+i+`">Delete</p></div>`;
	bb.innerHTML += rr
	Del(i)
}

let fortune = (e) => {
	let i = 0;
	console.log(e)
	let add = document.getElementById("add");
	add.addEventListener('click', event => {
		addArgs(i++, e);
	});
}

let Del = (i) => {
	let add = document.getElementById("del-"+i);
	add.addEventListener('click', event => {
		document.getElementById("delete-" + i).remove();
	})
}
