<script lang="ts">
    export let isDarkMode: boolean
    export let changeColorTheme = () => {}
</script>

<main>
    <button class={`mode-toggle ${isDarkMode ? "" : "toggle-right"}`} on:click={changeColorTheme}>
        <div class="slider">
    
        </div>
    </button>
</main>

<style lang="scss">
    $toggle-width: 45px;
    $padding:      5px;
    $slider-width: 15px;

    .mode-toggle{  
        display: flex;
        position: absolute;
        top: 20px;
        right: 20px;

        background-color: var(--twidder-blue);
        border-radius: 20px;
        border: none;

        width: $toggle-width;
        height: fit-content;
        padding: $padding $padding;

        .slider{
            background-color: white;
            width: $slider-width;
            height: $slider-width;
            border-radius: 50%;   
            
            transition: 0.3s;
            position: relative;
        }

        &:hover{
            cursor: pointer;
        }

        @media only screen and (min-width: 768px) {
            position: fixed;
        }
    }

    .toggle-right{
        .slider{
            transform: translateX(calc($toggle-width - $slider-width - ($padding * 2)));
        }
    }
</style>