<script setup>
import { Circle } from 'lucide-vue-next';

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
  primary: "text-center text-white bg-black hover:bg-slate-900 border-2 border-transparent",
  google: "text-white hover:text-black bg-blue-600 hover:bg-blue-600 border-2 border-transparent relative overflow-hidden shadow-2xl transition-all duration-500 before:absolute before:bottom-0 before:left-0 before:right-0 before:top-0 before:m-auto before:h-0 before:w-0 before:rounded-sm before:bg-white before:duration-500 before:ease-out hover:shadow-slate-600 hover:before:h-40 hover:before:w-full hover:before:bg-white",
  github: "text-white hover:text-black bg-black hover:bg-black border-2 border-transparent relative overflow-hidden shadow-2xl transition-all duration-500 before:absolute before:bottom-0 before:left-0 before:right-0 before:top-0 before:m-auto before:h-0 before:w-0 before:rounded-sm before:bg-white before:duration-500 before:ease-out hover:shadow-slate-600 hover:before:h-40 hover:before:w-full hover:before:bg-white",
  facebook: "text-white hover:text-black bg-[#3b5998] hover:bg-[#3b5998] border-2 border-transparent relative overflow-hidden shadow-2xl transition-all duration-500 before:absolute before:bottom-0 before:left-0 before:right-0 before:top-0 before:m-auto before:h-0 before:w-0 before:rounded-sm before:bg-white before:duration-500 before:ease-out hover:shadow-slate-600 hover:before:h-40 hover:before:w-full hover:before:bg-white",
  linkedin: "text-white hover:text-black bg-[#0e76a8] hover:bg-[#0e76a8] border-2 border-transparent relative overflow-hidden shadow-2xl transition-all duration-500 before:absolute before:bottom-0 before:left-0 before:right-0 before:top-0 before:m-auto before:h-0 before:w-0 before:rounded-sm before:bg-white before:duration-500 before:ease-out hover:shadow-slate-600 hover:before:h-40 hover:before:w-full hover:before:bg-white",
  apple: "text-white hover:text-blue-600 bg-black hover:bg-black border-2 border-transparent relative overflow-hidden shadow-2xl transition-all duration-500 before:absolute before:bottom-0 before:left-0 before:right-0 before:top-0 before:m-auto before:h-0 before:w-0 before:rounded-sm before:bg-white before:duration-500 before:ease-out hover:shadow-slate-600 hover:before:h-40 hover:before:w-full hover:before:bg-white",
  circle: "text-white hover:text-black relative overflow-hidden border bg-black border-black shadow-2xl transition-all duration-500 before:absolute before:bottom-0 before:left-0 before:right-0 before:top-0 before:m-auto before:h-0 before:w-0 before:rounded-sm before:bg-white before:duration-500 before:ease-out hover:text-blue-600 hover:shadow-indigo-600 hover:before:h-40 hover:before:w-full hover:before:opacity-80",
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
      'rounded transition focus-visible:ring-2 ring-offset-2 ring-gray-200 flex items-center gap-2 ',
      block && 'w-full',
      sizes[size],
      styles[styleName],
      className,
    ]"
  >
    <Icon v-if="icone != 'no'" :name="'uil:'+icone" class="w-6 h-6 " />
    <span class="flex-grow text-center"><slot /></span>
  </button>
</template>

e