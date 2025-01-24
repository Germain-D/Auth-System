<script setup>
const props = defineProps({
  size: {
    type: String,
    default: "md",
  },
  styleName: {
    type: String,
    default: "primary",
  },
  block: {
    type: Boolean,
    default: false,
  },
  className: {
    type: String,
    default: "",
  },
});

const sizes = {
  md: "px-5 py-2.5",
  lg: "px-6 py-3",
};

const styles = {
  outline: "text-center border-2 border-black hover:bg-black text-black hover:text-white",
  primary: "text-center bg-black text-white hover:bg-slate-900 border-2 border-transparent",
  google: "text-center bg-blue-600 text-white hover:bg-blue-600 border-2 border-transparent",
  github: "bg-black text-white hover:bg-black border-2 border-transparent flex items-center",
  facebook: "bg-[#3b5998] text-white hover:bg-[#3b5998] border-2 border-transparent",
  linkedin: "bg-[#0e76a8] text-white hover:bg-[#0e76a8] border-2 border-transparent",
  apple: "bg-black text-white hover:bg-black border-2 border-transparent",
};



let icone = 'no'
if (props.styleName === 'google') {
  icone = 'google'
} else if (props.styleName === 'github') {
  icone = 'github'
} else if (props.styleName === 'facebook') {
  icone = 'facebook'
} else if (props.styleName === 'linkedin') {
  icone = 'linkedin'
} else if (props.styleName === 'apple') {
  icone = 'apple'
}

</script>

<template>
  <button
    :class="[
      'rounded transition focus-visible:ring-2 ring-offset-2 ring-gray-200 flex items-center gap-2',
      block && 'w-full',
      sizes[size],
      styles[styleName],
      className,
    ]"
  >
    <Icon v-if="icone != 'no'" :name="'uil:'+icone" style="color: white" class="w-6 h-6" />
    <span class="flex-grow text-center"><slot /></span>
  </button>
</template>