const DEFALT_RESIZE_SCALE = "1.0";
const RESIZE_SCALE_INPUT = document.getElementById("resizeScaleId");
const RESIZE_INPUT_RANGE = document.getElementById("resizeInputRangeId");

RESIZE_SCALE_INPUT.innerText = DEFALT_RESIZE_SCALE;
RESIZE_INPUT_RANGE.value = DEFALT_RESIZE_SCALE;

const changeValue = input => {
    RESIZE_SCALE_INPUT.innerText = input.value === "1" ? "1.0" : input.value;
    RESIZE_INPUT_RANGE.value = input.value;
}

addEventListener("submit", () => {
    const resizeScale = RESIZE_INPUT_RANGE.value;
    const form = document.getElementById("formId");

    form.action = `${window.origin}?resize_scale=${resizeScale}`;
})

