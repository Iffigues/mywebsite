let addArgs = (i,e) => {
	let bb = document.getElementById("args");
	let rr = `<div id="delete-`+i+`"><input class="form-control" type="number" id=percent-"`+i+`" name="percent-`+i+`" min="0" max="100"><select class="form-control name="type-` + i + `">`		
	e.forEach((element) => {
		rr += `<option value="`+element+`">`+element+`</option>`
	} ) 
	rr +=  `</select><button type="button" class="btn btn-danger"id="del-`+i+`" onClick="Del(i)">Delete</button></div>`;
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
