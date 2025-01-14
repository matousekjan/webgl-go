package webgl

import (
	"errors"
	"syscall/js"
)

var (
	float32Array = js.Global().Get("Float32Array")
	uint8Array   = js.Global().Get("Uint8Array")
)

type ErrorNumber int
type ShaderType int
type BufferType int
type BufferUsage int
type Capacity int
type DepthFunc int
type Type int
type BufferMask int
type DrawMode int
type ProgramParameter int
type ShaderParameter int
type TextureType int
type PixelFormat int
type TextureParameter int
type TextureNumber int
type BlendFactor int
type BufferMode int
type TransformFeedbackTarget int
type SyncCondition int
type SyncFlushCommandBit int
type WaitSyncResult int
type FrameBufferType int

type Shader js.Value
type Program js.Value
type Location js.Value
type Buffer js.Value
type FrameBuffer js.Value
type Texture *js.Value
type TransformFeedback js.Value
type WebGLSync js.Value
type VertexArray js.Value

type WebGL struct {
	gl js.Value

	Canvas Canvas

	NO_ERROR, INVALID_ENUM, INVALID_VALUE, INVALID_OPERATION,
	INVALID_FRAMEBUFFER_OPERATION, OUT_OF_MEMORY, CONTEXT_LOST_WEBGL ErrorNumber

	FRAMEBUFFER, DRAW_FRAMEBUFFER, READ_FRAMEBUFFER                               FrameBufferType
	VERTEX_SHADER, FRAGMENT_SHADER                                                ShaderType
	ARRAY_BUFFER, ELEMENT_ARRAY_BUFFER, TRANSFORM_FEEDBACK_BUFFER, UNIFORM_BUFFER BufferType
	STATIC_DRAW, DYNAMIC_COPY, STREAM_READ                                        BufferUsage
	DEPTH_TEST, BLEND, POLYGON_OFFSET_FILL, RASTERIZER_DISCARD                    Capacity
	LEQUAL                                                                        DepthFunc
	FLOAT, UNSIGNED_BYTE, UNSIGNED_SHORT, UNSIGNED_INT                            Type
	COLOR_BUFFER_BIT, DEPTH_BUFFER_BIT, STENCIL_BUFFER_BIT                        BufferMask
	POINTS, LINE_STRIP, LINE_LOOP, LINES, TRIANGLE_STRIP, TRIANGLE_FAN, TRIANGLES DrawMode
	COMPILE_STATUS                                                                ShaderParameter
	LINK_STATUS, VALIDATE_STATUS                                                  ProgramParameter
	TEXTURE_2D                                                                    TextureType
	INTERLEAVED_ATTRIBS, SEPARATE_ATTRIBS                                         BufferMode
	TRANSFORM_FEEDBACK                                                            TransformFeedbackTarget
	SYNC_GPU_COMMANDS_COMPLETE                                                    SyncCondition
	SYNC_FLUSH_COMMANDS_BIT                                                       SyncFlushCommandBit
	MAX_CLIENT_WAIT_TIMEOUT_WEBGL                                                 int
	ALREADY_SIGNALED, TIMEOUT_EXPIRED, CONDITION_SATISFIED, WAIT_FAILED           WaitSyncResult

	RGBA PixelFormat

	TEXTURE_MIN_FILTER, TEXTURE_MAG_FILTER, TEXTURE_WRAP_S, TEXTURE_WRAP_T TextureParameter
	LINEAR, NEAREST, CLAMP_TO_EDGE                                         int

	TEXTURE0 TextureNumber

	ZERO, ONE, SRC_ALPHA, DST_ALPHA, ONE_MINUS_SRC_ALPHA, ONE_MINUS_DST_ALPHA, SRC_COLOR, DST_COLOR, SRC_ALPHA_SATURATE BlendFactor
}

func New(canvas js.Value) (*WebGL, error) {
	gl := canvas.Call("getContext", "webgl2")
	if gl.IsNull() {
		return nil, errors.New("WebGL is not supported")
	}
	return &WebGL{
		gl: gl,

		Canvas: Canvas(gl.Get("canvas")),

		NO_ERROR:                      ErrorNumber(gl.Get("NO_ERROR").Int()),
		INVALID_ENUM:                  ErrorNumber(gl.Get("INVALID_ENUM").Int()),
		INVALID_VALUE:                 ErrorNumber(gl.Get("INVALID_VALUE").Int()),
		INVALID_OPERATION:             ErrorNumber(gl.Get("INVALID_OPERATION").Int()),
		INVALID_FRAMEBUFFER_OPERATION: ErrorNumber(gl.Get("INVALID_FRAMEBUFFER_OPERATION").Int()),
		OUT_OF_MEMORY:                 ErrorNumber(gl.Get("OUT_OF_MEMORY").Int()),
		CONTEXT_LOST_WEBGL:            ErrorNumber(gl.Get("CONTEXT_LOST_WEBGL").Int()),

		VERTEX_SHADER:   ShaderType(gl.Get("VERTEX_SHADER").Int()),
		FRAGMENT_SHADER: ShaderType(gl.Get("FRAGMENT_SHADER").Int()),

		FRAMEBUFFER:      FrameBufferType(gl.Get("FRAMEBUFFER").Int()),
		DRAW_FRAMEBUFFER: FrameBufferType(gl.Get("DRAW_FRAMEBUFFER").Int()),
		READ_FRAMEBUFFER: FrameBufferType(gl.Get("READ_FRAMEBUFFER").Int()),

		ARRAY_BUFFER:              BufferType(gl.Get("ARRAY_BUFFER").Int()),
		ELEMENT_ARRAY_BUFFER:      BufferType(gl.Get("ELEMENT_ARRAY_BUFFER").Int()),
		TRANSFORM_FEEDBACK_BUFFER: BufferType(gl.Get("TRANSFORM_FEEDBACK_BUFFER").Int()),
		UNIFORM_BUFFER:            BufferType(gl.Get("UNIFORM_BUFFER").Int()),

		STATIC_DRAW:  BufferUsage(gl.Get("STATIC_DRAW").Int()),
		DYNAMIC_COPY: BufferUsage(gl.Get("DYNAMIC_COPY").Int()),
		STREAM_READ:  BufferUsage(gl.Get("STREAM_READ").Int()),

		DEPTH_TEST:          Capacity(gl.Get("DEPTH_TEST").Int()),
		BLEND:               Capacity(gl.Get("BLEND").Int()),
		POLYGON_OFFSET_FILL: Capacity(gl.Get("POLYGON_OFFSET_FILL").Int()),
		RASTERIZER_DISCARD:  Capacity(gl.Get("RASTERIZER_DISCARD").Int()),

		LEQUAL: DepthFunc(gl.Get("LEQUAL").Int()),

		FLOAT:          Type(gl.Get("FLOAT").Int()),
		UNSIGNED_BYTE:  Type(gl.Get("UNSIGNED_BYTE").Int()),
		UNSIGNED_SHORT: Type(gl.Get("UNSIGNED_SHORT").Int()),
		UNSIGNED_INT:   Type(gl.Get("UNSIGNED_INT").Int()),

		COLOR_BUFFER_BIT:   BufferMask(gl.Get("COLOR_BUFFER_BIT").Int()),
		DEPTH_BUFFER_BIT:   BufferMask(gl.Get("DEPTH_BUFFER_BIT").Int()),
		STENCIL_BUFFER_BIT: BufferMask(gl.Get("STENCIL_BUFFER_BIT").Int()),

		POINTS:         DrawMode(gl.Get("POINTS").Int()),
		LINE_STRIP:     DrawMode(gl.Get("LINE_STRIP").Int()),
		LINE_LOOP:      DrawMode(gl.Get("LINE_LOOP").Int()),
		LINES:          DrawMode(gl.Get("LINES").Int()),
		TRIANGLE_STRIP: DrawMode(gl.Get("TRIANGLE_STRIP").Int()),
		TRIANGLE_FAN:   DrawMode(gl.Get("TRIANGLE_FAN").Int()),
		TRIANGLES:      DrawMode(gl.Get("TRIANGLES").Int()),

		COMPILE_STATUS: ShaderParameter(gl.Get("COMPILE_STATUS").Int()),

		LINK_STATUS:     ProgramParameter(gl.Get("LINK_STATUS").Int()),
		VALIDATE_STATUS: ProgramParameter(gl.Get("VALIDATE_STATUS").Int()),

		TEXTURE_2D: TextureType(gl.Get("TEXTURE_2D").Int()),

		RGBA: PixelFormat(gl.Get("RGBA").Int()),

		TEXTURE_MIN_FILTER: TextureParameter(gl.Get("TEXTURE_MIN_FILTER").Int()),
		TEXTURE_MAG_FILTER: TextureParameter(gl.Get("TEXTURE_MAG_FILTER").Int()),
		TEXTURE_WRAP_S:     TextureParameter(gl.Get("TEXTURE_WRAP_S").Int()),
		TEXTURE_WRAP_T:     TextureParameter(gl.Get("TEXTURE_WRAP_T").Int()),

		LINEAR:        gl.Get("LINEAR").Int(),
		NEAREST:       gl.Get("NEAREST").Int(),
		CLAMP_TO_EDGE: gl.Get("CLAMP_TO_EDGE").Int(),

		TEXTURE0: TextureNumber(gl.Get("TEXTURE0").Int()),

		ZERO:                BlendFactor(gl.Get("ZERO").Int()),
		ONE:                 BlendFactor(gl.Get("ONE").Int()),
		SRC_ALPHA:           BlendFactor(gl.Get("SRC_ALPHA").Int()),
		DST_ALPHA:           BlendFactor(gl.Get("DST_ALPHA").Int()),
		ONE_MINUS_SRC_ALPHA: BlendFactor(gl.Get("ONE_MINUS_SRC_ALPHA").Int()),
		ONE_MINUS_DST_ALPHA: BlendFactor(gl.Get("ONE_MINUS_DST_ALPHA").Int()),
		SRC_COLOR:           BlendFactor(gl.Get("SRC_COLOR").Int()),
		DST_COLOR:           BlendFactor(gl.Get("DST_COLOR").Int()),
		SRC_ALPHA_SATURATE:  BlendFactor(gl.Get("SRC_ALPHA_SATURATE").Int()),

		INTERLEAVED_ATTRIBS: BufferMode(gl.Get("INTERLEAVED_ATTRIBS").Int()),
		SEPARATE_ATTRIBS:    BufferMode(gl.Get("SEPARATE_ATTRIBS").Int()),

		TRANSFORM_FEEDBACK: TransformFeedbackTarget(gl.Get("TRANSFORM_FEEDBACK").Int()),

		SYNC_GPU_COMMANDS_COMPLETE: SyncCondition(gl.Get("SYNC_GPU_COMMANDS_COMPLETE").Int()),

		SYNC_FLUSH_COMMANDS_BIT: SyncFlushCommandBit(gl.Get("SYNC_FLUSH_COMMANDS_BIT").Int()),

		MAX_CLIENT_WAIT_TIMEOUT_WEBGL: gl.Get("MAX_CLIENT_WAIT_TIMEOUT_WEBGL").Int(),

		ALREADY_SIGNALED:    WaitSyncResult(gl.Get("ALREADY_SIGNALED").Int()),
		TIMEOUT_EXPIRED:     WaitSyncResult(gl.Get("TIMEOUT_EXPIRED").Int()),
		CONDITION_SATISFIED: WaitSyncResult(gl.Get("CONDITION_SATISFIED").Int()),
		WAIT_FAILED:         WaitSyncResult(gl.Get("WAIT_FAILED").Int()),
	}, nil
}

func (gl *WebGL) GetError() error {
	e := ErrorNumber(gl.gl.Call("getError").Int())
	if e == gl.NO_ERROR {
		return nil
	}
	return &Error{
		Context: gl,
		Number:  e,
	}
}

func (gl *WebGL) CreateShader(t ShaderType) Shader {
	return Shader(gl.gl.Call("createShader", int(t)))
}

func (gl *WebGL) ShaderSource(s Shader, src string) {
	gl.gl.Call("shaderSource", js.Value(s), src)
}

func (gl *WebGL) CompileShader(s Shader) {
	gl.gl.Call("compileShader", js.Value(s))
}

func (gl *WebGL) GetShaderParameter(s Shader, param ShaderParameter) interface{} {
	v := gl.gl.Call("getShaderParameter", js.Value(s), int(param))
	switch param {
	case gl.COMPILE_STATUS:
		return v.Bool()
	}
	return nil
}

func (gl *WebGL) GetShaderInfoLog(s Shader) string {
	return gl.gl.Call("getShaderInfoLog", js.Value(s)).String()
}

func (gl *WebGL) CreateProgram() Program {
	return Program(gl.gl.Call("createProgram"))
}

func (gl *WebGL) AttachShader(p Program, s Shader) {
	gl.gl.Call("attachShader", js.Value(p), js.Value(s))
}

func (gl *WebGL) LinkProgram(p Program) {
	gl.gl.Call("linkProgram", js.Value(p))
}

func (gl *WebGL) GetProgramParameter(p Program, param ProgramParameter) interface{} {
	v := gl.gl.Call("getProgramParameter", js.Value(p), int(param))
	switch param {
	case gl.LINK_STATUS, gl.VALIDATE_STATUS:
		return v.Bool()
	}
	return nil
}

func (gl *WebGL) GetProgramInfoLog(p Program) string {
	return gl.gl.Call("getProgramInfoLog", js.Value(p)).String()
}

func (gl *WebGL) UseProgram(p Program) {
	gl.gl.Call("useProgram", js.Value(p))
}

func (gl *WebGL) GetAttribLocation(p Program, name string) int {
	return gl.gl.Call("getAttribLocation", js.Value(p), name).Int()
}

func (gl *WebGL) GetUniformLocation(p Program, name string) Location {
	return Location(gl.gl.Call("getUniformLocation", js.Value(p), name))
}

func (gl *WebGL) CreateVertexArray() VertexArray {
	return VertexArray(gl.gl.Call("createVertexArray"))
}

func (gl *WebGL) BindVertexArray(vertArr VertexArray) {
	gl.gl.Call("bindVertexArray", js.Value(vertArr))
}

func (gl *WebGL) BindFrameBuffer(ft FrameBufferType, fBuf FrameBuffer) {
	gl.gl.Call("bindFramebuffer", int(ft), js.Value(fBuf))
}

func (gl *WebGL) CreateBuffer() Buffer {
	return Buffer(gl.gl.Call("createBuffer"))
}

func (gl *WebGL) BindBuffer(t BufferType, buf Buffer) {
	gl.gl.Call("bindBuffer", int(t), js.Value(buf))
}

func (gl *WebGL) BindBufferBase(target BufferType, index int, buf Buffer) {
	gl.gl.Call("bindBufferBase", int(target), index, js.Value(buf))
}

func (gl *WebGL) BufferData(t BufferType, data BufferData, usage BufferUsage) {
	bin := data.Bytes()
	dataJS := uint8Array.New(len(bin))
	js.CopyBytesToJS(dataJS, bin)
	gl.gl.Call("bufferData", int(t), dataJS, int(usage))
}

func (gl *WebGL) BufferData_JS(t BufferType, data js.Value, usage BufferUsage) {
	gl.gl.Call("bufferData", int(t), data, int(usage))
}

func (gl *WebGL) GetBufferSubData(t BufferType, srcOffset int, view js.Value, dstOffset, length int) {
	gl.gl.Call("getBufferSubData", int(t), srcOffset, view, dstOffset, length)
}

func (gl *WebGL) ClearColor(r, g, b, a float32) {
	gl.gl.Call("clearColor", r, g, b, a)
}

func (gl *WebGL) ClearDepth(d float32) {
	gl.gl.Call("clearDepth", d)
}

func (gl *WebGL) Enable(c Capacity) {
	gl.gl.Call("enable", int(c))
}

func (gl *WebGL) Disable(c Capacity) {
	gl.gl.Call("disable", int(c))
}

func (gl *WebGL) IsEnabled(c Capacity) bool {
	return gl.gl.Call("isEnabled", int(c)).Bool()
}

func (gl *WebGL) DepthFunc(f DepthFunc) {
	gl.gl.Call("depthFunc", int(f))
}

func (gl *WebGL) VertexAttribPointer(i, size int, typ Type, normalized bool, stride, offset int) {
	gl.gl.Call("vertexAttribPointer", i, size, int(typ), normalized, stride, offset)
}

func (gl *WebGL) VertexAttribIPointer(i, size int, typ Type, stride, offset int) {
	gl.gl.Call("vertexAttribIPointer", i, size, int(typ), stride, offset)
}

func (gl *WebGL) VertexAttribDivisor(index int, divisor int) {
	gl.gl.Call("vertexAttribDivisor", index, divisor)
}

func (gl *WebGL) EnableVertexAttribArray(i int) {
	gl.gl.Call("enableVertexAttribArray", i)
}

func (gl *WebGL) DisableVertexAttribArray(i int) {
	gl.gl.Call("disableVertexAttribArray", i)
}

func (gl *WebGL) UniformMatrix4fv(loc Location, transpose bool, mat Mat4) {
	f := mat.Floats()
	matJS := float32Array.Call("of",
		f[0], f[1], f[2], f[3],
		f[4], f[5], f[6], f[7],
		f[8], f[9], f[10], f[11],
		f[12], f[13], f[14], f[15],
	)
	gl.gl.Call("uniformMatrix4fv", js.Value(loc), transpose, matJS)
}

func (gl *WebGL) Uniform3fv(loc Location, v Vec3) {
	f := v.Floats()
	vecJS := float32Array.Call("of", f[0], f[1], f[2])
	gl.gl.Call("uniform3fv", js.Value(loc), vecJS)
}

func (gl *WebGL) Uniform1i(loc Location, i int) {
	gl.gl.Call("uniform1i", js.Value(loc), i)
}

func (gl *WebGL) Uniform1ui(loc Location, i uint32) {
	gl.gl.Call("uniform1ui", js.Value(loc), i)
}

func (gl *WebGL) Uniform1f(loc Location, i float32) {
	gl.gl.Call("uniform1f", js.Value(loc), i)
}

func (gl *WebGL) Clear(mask BufferMask) {
	gl.gl.Call("clear", int(mask))
}

func (gl *WebGL) DrawArrays(mode DrawMode, i, n int) {
	gl.gl.Call("drawArrays", int(mode), i, n)
}

func (gl *WebGL) DrawElements(mode DrawMode, count int, typ Type, offset int) {
	gl.gl.Call("drawElements", int(mode), count, int(typ), offset)
}

func (gl *WebGL) DrawElementsInstanced(mode DrawMode, count int, typ Type, offset int, instanceCount int) {
	gl.gl.Call("drawElementsInstanced", int(mode), count, int(typ), offset, instanceCount)
}

func (gl *WebGL) Viewport(x1, y1, x2, y2 int) {
	gl.gl.Call("viewport", x1, y1, x2, y2)
}

func (gl *WebGL) CreateTexture() Texture {
	tex := gl.gl.Call("createTexture")
	return Texture(&tex)
}

func (gl *WebGL) BindTexture(texType TextureType, tex Texture) {
	if tex == nil {
		gl.gl.Call("bindTexture", int(texType), nil)
		return
	}
	gl.gl.Call("bindTexture", int(texType), js.Value(*tex))
}

func (gl *WebGL) TexImage2D(texType TextureType, level int, internalFmt, fmt PixelFormat, typ Type, img interface{}) {
	gl.gl.Call("texImage2D", int(texType), level, int(internalFmt), int(fmt), int(typ), img)
}

func (gl *WebGL) TexParameteri(texType TextureType, param TextureParameter, val interface{}) {
	gl.gl.Call("texParameteri", int(texType), int(param), val)
}

func (gl *WebGL) ActiveTexture(i TextureNumber) {
	gl.gl.Call("activeTexture", int(i))
}

func (gl *WebGL) BlendFunc(s, d BlendFactor) {
	gl.gl.Call("blendFunc", int(s), int(d))
}

func (gl *WebGL) TransformFeedbackVaryings(p Program, varyings []string, m BufferMode) {
	var v []interface{}
	for _, vs := range varyings {
		v = append(v, vs)
	}
	gl.gl.Call("transformFeedbackVaryings", js.Value(p), js.ValueOf(v), int(m))
}

func (gl *WebGL) BeginTransformFeedback(mode DrawMode) {
	gl.gl.Call("beginTransformFeedback", int(mode))
}

func (gl *WebGL) EndTransformFeedback() {
	gl.gl.Call("endTransformFeedback")
}

func (gl *WebGL) CreateTransformFeedback() TransformFeedback {
	return TransformFeedback(gl.gl.Call("createTransformFeedback"))
}

func (gl *WebGL) BindTransformFeedback(target TransformFeedbackTarget, fb TransformFeedback) {
	gl.gl.Call("bindTransformFeedback", int(target), js.Value(fb))
}

func (gl *WebGL) FenceSync(c SyncCondition, flags int) WebGLSync {
	return WebGLSync(gl.gl.Call("fenceSync", int(c), flags))
}

func (gl *WebGL) ClientWaitSync(sync WebGLSync, flags SyncFlushCommandBit, timeout int) WaitSyncResult {
	return WaitSyncResult(gl.gl.Call("clientWaitSync", js.Value(sync), int(flags), timeout).Int())
}

func (gl *WebGL) DeleteSync(sync WebGLSync) {
	gl.gl.Call("deleteSync", js.Value(sync))
}

func (gl *WebGL) Flush() {
	gl.gl.Call("flush")
}

func (gl *WebGL) Finish() {
	gl.gl.Call("finish")
}

func (gl *WebGL) IsContextLost() bool {
	return gl.gl.Call("isContextLost").Bool()
}

func (gl *WebGL) PolygonOffset(factor, units float32) {
	gl.gl.Call("polygonOffset", factor, units)
}

func (gl *WebGL) JS() js.Value {
	return gl.gl
}

func (gl *WebGL) GetExtension(name string) (js.Value, bool) {
	ret := gl.gl.Call("getExtension", name)
	if ret.IsNull() {
		return js.Value{}, false
	}
	return ret, true
}

func (gl *WebGL) GetParameter(id int) js.Value {
	return gl.gl.Call("getParameter", id)
}
