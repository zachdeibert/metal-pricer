"use strict";

function r(db) {
    var initialFields = [
        "type",
        "shape"
    ];

    var fields = [];
    db.forEach(el => {
        Object.getOwnPropertyNames(el).forEach(prop => {
            if (typeof(el[prop]) === "string") {
                if (!fields.includes(prop)) {
                    fields.push(prop);
                }
            }
        })
    });
    initialFields.forEach(field => {
        var idx = fields.indexOf(field);
        if (idx < 0) {
            console.error(`Unknown field ${field}`);
        } else {
            fields.splice(idx, 1);
        }
    });

    var dom = {
        "shapeOptions": document.getElementById("shapeOptions"),
        "sizeOptions": document.getElementById("sizeOptions"),
        "unitResult": document.getElementById("unitResult"),
        "unit": document.getElementById("unit"),
        "unitPrice": document.getElementById("unitPrice"),
        "sizeResult": document.getElementById("sizeResult"),
        "size": document.getElementById("size"),
        "sizePrice": document.getElementById("sizePrice"),
    };
    var search = {};
    var size = {};

    function removeAllChildren(node) {
        while (node.firstChild) {
            node.removeChild(node.firstChild);
        }
    }

    function getPossibleValues(field) {
        var possible = [];
        db.forEach(el => {
            var match = true;
            Object.getOwnPropertyNames(search).forEach(f => {
                if (el[f] !== search[f]) {
                    match = false;
                }
            });
            if (match && !possible.includes(el[field])) {
                possible.push(el[field]);
            }
        });
        return possible;
    }

    function getInitialConstrainedPossibleValues(field) {
        var tmp = search;
        search = {};
        initialFields.forEach(f => {
            if (tmp[f]) {
                search[f] = tmp[f];
            }
        });
        var res = getPossibleValues(field);
        search = tmp;
        return res;
    }

    function getAllPossibleValues(field) {
        var tmp = search;
        search = {};
        var res = getPossibleValues(field);
        search = tmp;
        return res;
    }

    function dbLookup() {
        var res = db.filter(el => {
            var valid = true;
            Object.getOwnPropertyNames(search).forEach(field => {
                if (el[field] !== search[field]) {
                    valid = false;
                }
            });
            return valid;
        });
        if (res.length !== 1) {
            console.error("dbLookup called with non-specific search");
        }
        return res[0];
    }

    function renderSelectBox(container, field, all) {
        var options = (all === true ? getAllPossibleValues : getPossibleValues)(field);
        var div = document.createElement("div");
        var label = document.createElement("label");
        label.innerText = `${field}:`;
        div.appendChild(label);
        var select = document.createElement("select");
        select.addEventListener("change", () => {
            search[field] = select.value;
            initialFields.forEach(f => {
                if (f !== field && getInitialConstrainedPossibleValues(f).length === 0) {
                    delete search[f];
                }
                if (f === field) {
                    fields.forEach(f => {
                        delete search[f];
                    });
                    size = {};
                }
            });
            updateShapeOptions();
        });
        if (!search[field]) {
            var option = document.createElement("option");
            select.appendChild(option);
        }
        options.forEach(opt => {
            var option = document.createElement("option");
            option.value = opt;
            option.innerText = opt;
            option.selected = search[field] == opt;
            select.appendChild(option);
        });
        div.appendChild(select);
        container.appendChild(div);
    }

    function renderTextBox(container, name) {
        var div = document.createElement("div");
        var label = document.createElement("label");
        label.innerText = `${name}:`;
        div.appendChild(label);
        var input = document.createElement("input");
        input.addEventListener("change", () => {
            size[name] = parseInt(input.value);
            updateSize();
        });
        input.type = "number";
        if (size[name]) {
            input.value = size[name];
        }
        div.appendChild(input);
        var span = document.createElement("span");
        span.innerText = "ft";
        div.appendChild(span);
        container.appendChild(div);
    }

    function updateSize() {
        var match = dbLookup();
        if (match.isSquare) {
            if (!(dom.sizeResult.hidden = !size.length || !size.width)) {
                dom.size.innerText = `${size.length}ft x ${size.width}ft`;
                dom.sizePrice.innerText = (match.price * size.length * size.width).toFixed(2);
            }
        } else {
            if (!(dom.sizeResult.hidden = !size.length)) {
                dom.size.innerText = `${size.length}ft`;
                dom.sizePrice.innerText = (match.price * size.length).toFixed(2);
            }
        }
    }

    function updateShapeOptions() {
        var shapeContainer = document.createElement("div");
        var hasInitial = true;
        var hasAll = true;
        initialFields.forEach(field => {
            if (!search[field]) {
                hasInitial = false;
                hasAll = false;
            }
            renderSelectBox(shapeContainer, field, true);
        });
        if (hasInitial) {
            fields.forEach(field => {
                if (search[field] || getPossibleValues(field).length > 1) {
                    renderSelectBox(shapeContainer, field);
                    if (!search[field]) {
                        hasAll = false;
                    }
                }
            });
        }
        removeAllChildren(dom.shapeOptions);
        dom.shapeOptions.appendChild(shapeContainer);
        var sizeContainer = document.createElement("div");
        if (hasAll) {
            renderTextBox(sizeContainer, "length");
            var match = dbLookup();
            if (match.isSquare) {
                renderTextBox(sizeContainer, "width");
            }
            dom.unit.innerText = match.isSquare ? "square foot" : "foot";
            dom.unitPrice.innerText = match.price.toFixed(2);
            updateSize();
        } else {
            dom.sizeResult.hidden = true;
        }
        removeAllChildren(dom.sizeOptions);
        dom.sizeOptions.appendChild(sizeContainer);
        dom.unitResult.hidden = !hasAll;
    }
    updateShapeOptions();
}
