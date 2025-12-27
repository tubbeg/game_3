components {
  id: "card"
  component: "/card/card.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"back_1\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/card/card.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 2.0
    y: 2.0
  }
}
