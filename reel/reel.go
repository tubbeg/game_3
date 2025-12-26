components {
  id: "reel"
  component: "/reel/reel.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"beta\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/reel/reel.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 2.0
    y: 2.0
  }
}
