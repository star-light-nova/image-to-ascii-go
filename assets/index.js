const DEFALT_RESIZE_SCALE = "1.0";
const RESIZE_SCALE_INPUT = document.getElementById("resizeScaleId");
const RESIZE_INPUT_RANGE = document.getElementById("resizeInputRangeId");

RESIZE_SCALE_INPUT.innerText = DEFALT_RESIZE_SCALE;
RESIZE_INPUT_RANGE.value = DEFALT_RESIZE_SCALE;

function changeValue(input) {
    RESIZE_SCALE_INPUT.innerText = input.value === "1" ? "1.0" : input.value;
    RESIZE_INPUT_RANGE.value = input.value;
}

function submit() {
    const resizeScale = RESIZE_SCALE_INPUT.value;

    const form = document.getElementById("formId");
    form.action = `{{ .URL }}?resize_scale=${resizeScale}`;
}
