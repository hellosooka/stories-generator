import File3 from './File3.vue'
import { Meta, Story } from "@storybook/vue3";
import { defineComponent } from "vue";

export default {
  title: "testFileTree/File3",
  component: File3,
  argTypes: {
  },
  parameters: {
    docs: {
      source: {
        code: [
          "import { File3 } from '@estelink/testFileTree'\n\n",
          "<File3 />",
        ].join("\n"),
      },
    },
  },
};

const Template: Story = (args) =>
  defineComponent({
    components: { File3 },
    setup: () => ({ args }),
    template: `<File3 v-bind="args" />`,
  });

export const Default = Template.bind({});
